package go2cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
)

//cache channel ref
var cacheChannel *CacheChannel
var synOnce sync.Once

//cache channel single object
//init go2cache
func GetCacheChannel() *CacheChannel {
	if cacheChannel == nil {
		synOnce.Do(func() {
			//init cache channel
			cacheChannel = &CacheChannel{}
			_ = cacheChannel.initCacheChannel()
		})
	}
	return cacheChannel
}

//缓存操作入口
type CacheChannel struct {
	//mm provider
	mmp *MemoryProvider
	//redis provider
	rdp *RedisProvider
	//redis get mutex
	readDataMutex sync.Mutex
	//redis pub and sub
	psb *PubSub
}

//init go2cache
//path for config
func (c *CacheChannel) initCacheChannel() error {
	//get go2cache config
	cfg := GetConfig()
	//init memory cache
	if len(cfg.Regions) != 0 {
		//init MemoryProvider and RedisProvider
		c.mmp = new(MemoryProvider)
		c.rdp = new(RedisProvider)
		//init redis cache
		c.rdp.InitRedisClient()
		//init psc with redis
		if cfg.Psb {
			c.psb = &PubSub{
				Client:  c.rdp.redisClient,
				Channel: cfg.Channel,
				Region:  cfg.RedisNameSpace,
			}
			c.psb.Subscribe()
		}
		//redis name space
		c.rdp.redisNameSpace = cfg.RedisNameSpace
		for _, region := range cfg.Regions {
			_, _ = c.mmp.BuildCache(region.Name) //region size no limit
			_, _ = c.rdp.BuildCache(region.Name)
		}
	}
	return nil
}

//Get mm cache by region
func (c *CacheChannel) GetMemoryCache(region string) *MemoryCache {
	memoryCache, _ := c.mmp.BuildCache(region)
	return memoryCache.(*MemoryCache)
}

//Get redis cache by region
func (c *CacheChannel) GetRedisCache(region string) *RedisCache {
	redisCache, _ := c.rdp.BuildCache(region)
	return redisCache.(*RedisCache)
}

//del data form region cache
func (c *CacheChannel) Delete(region, key string) {
	//level1
	_ = c.GetMemoryCache(region).Delete(key)
	//level2
	c.GetRedisCache(region).Del(key)
	//notify
	c.SendEvictCmd(region, key)
}

//base protobuf struck ,read from level1 cache
func (c *CacheChannel) GetLevel1(region, key string) interface{} {
	memoryCache, _ := c.mmp.BuildCache(region)
	return memoryCache.(*MemoryCache).Get(key)
}

//base protobuf struck ,read from level2 cache
//may be return nil
//set protobuf
func (c *CacheChannel) GetProtoBufLevel2(region, key string, message proto.Message) interface{} {
	//must be mutex ,in case of more i/o in redis cache
	c.readDataMutex.Lock()
	defer c.readDataMutex.Unlock()

	memoryCache, _ := c.mmp.BuildCache(region)
	value := memoryCache.(*MemoryCache).Get(key) //SetProtoBuf pb struct
	if value != nil {
		return value
	}

	//unhit level1 ,try to read form level2 cache
	redisCache, _ := c.rdp.BuildCache(region)
	bytes, err := redisCache.(*RedisCache).GetBytes(key)
	if err != nil {
		log.Printf("GetProtoBufLevel2 error:%s  key:%s", err, key)
		return nil
	} else if bytes != nil {
		_ = proto.Unmarshal(bytes, message)
		_ = memoryCache.(*MemoryCache).Put(key, message)
		return message
	} else {
		return nil
	}
}

//read from level2 cache
func (c *CacheChannel) GetBytesLevel2(region, key string) ([]byte, error) {
	redisCache, _ := c.rdp.BuildCache(region)
	return redisCache.(*RedisCache).GetBytes(key)
}

//base protobuf .set into cache
//pusub notify other node (x2cache) evict level1 cache
func (c *CacheChannel) SetProtoBuf(region, key string, message proto.Message) {

	//set into mem cache
	//pb struck
	//memoryCache, _ := c.mmp.BuildCache(region)
	//_ = memoryCache.(*MemoryCache).Put(key, message)

	//set redis cache
	redisCache, _ := c.rdp.BuildCache(region)
	bytes, marshalError := proto.Marshal(message)
	if marshalError != nil {
		log.Printf("proto Marshal error:%s", marshalError)
		return
	}
	redisCache.(*RedisCache).Set(key, bytes)

	//remove other service level-1 cache
	c.SendEvictCmd(region, key)
}

//set into cache
//pusub notify other node (x2cache) evict level1 cache
func (c *CacheChannel) Set(region, key string, value interface{}) {
	memoryCache, _ := c.mmp.BuildCache(region)
	_ = memoryCache.(*MemoryCache).Put(key, value)
	redisCache, _ := c.rdp.BuildCache(region)
	redisCache.(*RedisCache).Set(key, value)
	//clear leve1 cache
	c.SendEvictCmd(region, key)
}

//send evict cmd to nodes
func (c *CacheChannel) SendEvictCmd(region, key string) {
	//log.Println("发送清理指令:", region+"@"+key)
	c.psb.SendEvictCmd(region, key)
}

//evict level1 cache
func (c *CacheChannel) Evict(region string, keys []string) {
	cache, _ := c.mmp.BuildCache(region)
	for _, key := range keys {
		_ = cache.(*MemoryCache).Delete(key)
	}
}

//get all regions
func (c *CacheChannel) GetRegions() []*Region {
	return GetConfig().Regions
}

//mm provider
func (c *CacheChannel) GetMemoryProvider() *MemoryProvider {
	return c.mmp
}

//redis provider
func (c *CacheChannel) GetRedisProvider() *RedisProvider {
	return c.rdp
}

//
type PubSub struct {
	Client  *redis.Client
	Channel string
	Region  string
}

//j2cache  发布 订阅消息模块 封装
const (
	OptJoin     = iota + 1 // 加入集群
	OptEvictKey            // 删除集群
	OptClearKey            // 清理缓存
	OptQuit                // 退出集群环境
)

//j2cache command
type Command struct {
	Region   string   `json:"region"`
	Operator int      `json:"operator"`
	Keys     []string `json:"keys"`
	Src      int      `json:"src"`
}

//发送清楚缓存的广播命令
func (p *PubSub) SendEvictCmd(region string, keys ...string) {
	data, _ := json.Marshal(&Command{Region: region, Keys: keys, Operator: OptEvictKey})
	intCmd := p.Client.Publish(p.Channel, data)
	e := intCmd.Err()
	if e != nil {
		log.Printf("error in pubish , info:%s", e.Error())
	}
}

//发送清除缓存的广播命令
func (p *PubSub) SendClearCmd(region string) {
	data, _ := json.Marshal(&Command{Region: region, Keys: nil, Operator: OptClearKey})
	intCmd := p.Client.Publish(p.Channel, data)
	e := intCmd.Err()
	if e != nil {
		log.Printf("error in pubish , info:%s", e.Error())
	}
}

//初始化订阅
//基于go-redis,可实现断线自动重连
func (p *PubSub) Subscribe() {
	psc := p.Client.Subscribe(p.Channel)
	//open  , in case of main thread blocking!!!
	go func() {
		for {
			//blocking until receive  data
			//if disconnect ,will auto retry!!!
			msg, _ := psc.Receive()
			switch msg.(type) {
			case *redis.Message:
				var cmd Command
				message := msg.(*redis.Message)
				e := json.Unmarshal([]byte(message.Payload), &cmd)
				if e != nil {
					log.Printf("command unmarshl json error:%s", e)
				}
				if cmd.Operator == OptEvictKey { //删除一级缓存数据
					cacheChannel.Evict(cmd.Region, cmd.Keys)
				} else if cmd.Operator == OptClearKey { //  清除缓存
					log.Printf("clear cache  key :%s region:%s", cmd.Keys, cmd.Region)
				} else if cmd.Operator == OptJoin { // 节点加入
					log.Printf("node-%d join cluster ", cmd.Src)
				} else if cmd.Operator == OptQuit { //节点离开
					log.Printf("node-%d quit cluster ", cmd.Src)
				}
			case *redis.Subscription:
				message := msg.(*redis.Subscription)
				log.Printf("subscription  channel :%s  %s  count:%d\n", message.Channel, message.Kind, message.Count)
			}
		}
	}()
}
