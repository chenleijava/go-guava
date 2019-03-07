package go2cache

import "C"
import (
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
			c.psb = &PubSub{Client: c.rdp.redisClient, Channel: cfg.Channel, cacheChannel: c, Region: cfg.RedisNameSpace}
			c.psb.Subscribe()
		}
		//redis name space
		c.rdp.redisNameSpace = cfg.RedisNameSpace
		for _, region := range cfg.Regions {
			c.mmp.BuildCache(region.Name) //region size no limit
			c.rdp.BuildCache(region.Name)
		}
	}
	return nil
}

//Get redis cache by region
func (c *CacheChannel) GetRedisCache(region string) *RedisCache {
	redisCache, _ := c.rdp.BuildCache(region)
	return redisCache.(*RedisCache)
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
		proto.Unmarshal(bytes, message)
		memoryCache.(*MemoryCache).Put(key, message)
		return message
	} else {
		return nil
	}
}

//read from level2 cache
func (c *CacheChannel) GetBytesLevel2(region, key string) ([]byte) {
	redisCache, _ := c.rdp.BuildCache(region)
	dd, _ := redisCache.(*RedisCache).GetBytes(key)
	return dd
}

//base protobuf .set into cache
//pusub notify other node (x2cache) evict level1 cache
func (c *CacheChannel) SetProtoBuf(region, key string, message proto.Message) {

	//set into mem cache
	//pb struck
	memoryCache, _ := c.mmp.BuildCache(region)
	_ = memoryCache.(*MemoryCache).Put(key, message)

	//set redis cache
	redisCache, _ := c.rdp.BuildCache(region)
	bytes, marshalError := proto.Marshal(message)
	if marshalError != nil {
		log.Printf("proto Marshal error:%s", marshalError)
		return
	}
	redisCache.(*RedisCache).Set(key, bytes)

	//remove other service level-1 cache
	c.sendEvictCmd(region, key)
}

//set into cache
//pusub notify other node (x2cache) evict level1 cache
func (c *CacheChannel) Set(region, key string, value interface{}) {
	memoryCache, _ := c.mmp.BuildCache(region)
	memoryCache.(*MemoryCache).Put(key, value)
	redisCache, _ := c.rdp.BuildCache(region)
	redisCache.(*RedisCache).Set(key, value)
	//clear leve1 cache
	c.sendEvictCmd(region, key)
}

//send evict cmd to nodes
func (c *CacheChannel) sendEvictCmd(region, key string) {
	//log.Println("发送清理指令:", region+"@"+key)
	c.psb.SendEvictCmd(region, key)
}

//evict level1 cache
func (c *CacheChannel) Evict(region string, keys []string) {
	cache, _ := c.mmp.BuildCache(region)
	for _, key := range keys {
		cache.(*MemoryCache).Delete(key)
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
