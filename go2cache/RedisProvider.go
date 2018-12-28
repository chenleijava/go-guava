package go2cache

import (
	"github.com/go-redis/redis"
	"runtime"
	"sync"
)

//基于redis缓存提供者
type RedisProvider struct {
	//redisCacheMap lock
	mapLock sync.Mutex
	//redis cache
	redisCacheMap map[string]*RedisCache
	// pool
	redisClient *redis.Client
	//redis name space
	redisNameSpace string
}

//init redis cache
func (p *RedisProvider) InitRedisClient() {
	p.mapLock.Lock()
	defer p.mapLock.Unlock()
	p.redisCacheMap = make(map[string]*RedisCache)
	p.redisClient = RedisClient() //only one,init
}

// build cache
func (p *RedisProvider) BuildCache(region string) (interface{}, error) {
	p.mapLock.Lock()
	defer p.mapLock.Unlock()
	region = p.redisNameSpace + ":" + region
	cache := p.redisCacheMap[region]
	if cache == nil {
		cache = &RedisCache{
			redisClient: p.redisClient,
			region:      region,
		}
		p.redisCacheMap[region] = cache
	}
	return cache, nil
}

// 缓存 等级
func (p *RedisProvider) Level() int {
	return Level2
}

//region name default  go2cache_redis
func (p *RedisProvider) Name() string {
	return "go2cache_redis_provider"
}

//获取region 列表
func (p *RedisProvider) GetRegions() []Region {
	return nil
}

//----------------------
//init redis client flg
var onceSingleRedisClient sync.Once
//go-redis client
var client *redis.Client
//get redis client
//disconnect will try to connect
func RedisClient() *redis.Client {
	if client == nil {
		onceSingleRedisClient.Do(func() {
			var cfg = GetConfig()
			//default  connections and minIdle equal  cpu number*10,db index 0;
			size := runtime.NumCPU() * 10
			options := &redis.Options{
				Addr:         cfg.Addr,
				Password:     cfg.Password, // no password set
				DB:           cfg.Db,       // use default DB
				MinIdleConns: size,
				PoolSize:     size,
			}
			if cfg.MinIdleConns != 0 {
				options.MinIdleConns = cfg.MinIdleConns
			}
			if cfg.PoolSize != 0 {
				options.PoolSize = cfg.PoolSize
			}
			if cfg.Db != 0 {
				options.DB = cfg.Db
			}
			client = redis.NewClient(options)
			//init done
		})
	}
	return client
}
