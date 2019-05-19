package go2cache

import (
	"github.com/go-redis/redis"
)

type RedisCache struct {
	redisClient *redis.Client
	region      string // region   -->  redis_name_space+":"+region
}

//get redis client
func (cache *RedisCache) RedisClient() *redis.Client {
	return cache.redisClient
}

//build region redis cache
//must call this function!!!
func (cache *RedisCache) BuildKey(key string) string {
	return cache.region + ":" + key
}

//redis 缓存中获取byte[]
func (cache *RedisCache) GetBytes(key string) ([]byte, error) {
	return cache.redisClient.Get(cache.BuildKey(key)).Bytes()
}

//存储数据到当前cache中
//timeout 对象有效期 0 永不过期
func (cache *RedisCache) Set(key string, value interface{}) interface{} {
	return cache.redisClient.Set(cache.BuildKey(key), value, 0).Val()
}

//删除缓存数据
func (cache *RedisCache) Del(key string) interface{} {
	return cache.redisClient.Del(cache.BuildKey(key)).Val()
}

//检查当前key 是否存在
func (cache *RedisCache) IsExist(key string) bool {
	return cache.redisClient.Exists(cache.BuildKey(key)).Val() == 1
}

//计数 +1
func (cache *RedisCache) Incr(key string) interface{} {
	return cache.redisClient.Incr(cache.BuildKey(key)).Val()
}

//获取 key 对应的值
func (cache *RedisCache) Get(key string) string {
	return cache.redisClient.Get(cache.BuildKey(key)).Val()
}

//hincryBy 基于hash 计数
func (cache *RedisCache) HincrBy(key, filed string, value int) interface{} {
	return cache.redisClient.HIncrBy(cache.BuildKey(key), filed, int64(value)).Val()
}

//HSET
func (cache *RedisCache) Hset(key, filed string, value interface{}) interface{} {
	return cache.redisClient.HSet(cache.BuildKey(key), filed, value).Val()
}

//HGET
func (cache *RedisCache) Hget(key, filed string) string {
	return cache.redisClient.HGet(cache.BuildKey(key), filed).Val()
}

//HDEL
func (cache *RedisCache) Hdel(key, filed string) int64 {
	return cache.redisClient.HDel(cache.BuildKey(key), filed).Val()
}

//HGETALL
//数据类型转换-需求的类型
func (cache *RedisCache) HgetAllStringMap(key string) map[string]string {
	return cache.redisClient.HGetAll(cache.BuildKey(key)).Val()
}

//HMGET
//return  map object ,  value maybe nil
func (cache *RedisCache) HMGet(key string, fields ...string) []interface{} {
	return cache.redisClient.HMGet(cache.BuildKey(key), fields...).Val()
}

//HMSet
func (cache *RedisCache) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	return cache.redisClient.HMSet(cache.BuildKey(key), fields)
}

//HLen
func (cache *RedisCache) Hlen(key string) int {
	return int(cache.redisClient.HLen(cache.BuildKey(key)).Val())
}

//sadd
func (cache *RedisCache) SAdd(key string, members ...interface{}) (int64, error) {
	return cache.redisClient.SAdd(cache.BuildKey(key), members...).Result()
}

//spopn
func (cache *RedisCache) SPopN(key string, count int64) ([]string, error) {
	return cache.redisClient.SPopN(cache.BuildKey(key), count).Result()
}

//smembers
func (cache *RedisCache) SMembers(key string) ([]string, error) {
	return cache.redisClient.SMembers(cache.BuildKey(key)).Result()
}

//sremove
func (cache *RedisCache) SRem(key string, members ...interface{}) (int64, error) {
	return cache.redisClient.SRem(cache.BuildKey(key), members...).Result()
}
