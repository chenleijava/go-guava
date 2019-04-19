package go2cache

import (
	"github.com/chenleijava/go-guava"
	"github.com/chenleijava/go-guava/zlog"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"log"
	"strconv"
	"testing"
	"time"
)

var logger = zlog.NewLog2Console()

//test go2cache
func TestGetCacheChannel(t *testing.T) {
	cacheChannel := GetCacheChannel()
	regin := "login_log_region"
	cache := cacheChannel.GetRedisCache("login_log_region")

	c := cache.redisClient
	kk := cache.BuildKey("1")
	c.HMSet(kk, map[string]interface{}{
		"like_num": 1,
	})

	s := c.HMGet(cache.BuildKey("3"), "like_num","x")
	if len(s.Val()) != 0 {
		num := s.Val()[0]
		t, _ := strconv.ParseInt(num.(string), 10, 32)
		log.Printf("%d", t)
	}

	for {
		select {
		case <-time.Tick(time.Second * 1):

			cacheChannel.SendEvictCmd(regin, "123")
		}
	}

	var key = time.Now().Format("2006-01-02 15:04:05")
	var filed = key
	v := cache.HincrBy(key, filed, 1)
	log.Printf("hincyBy v:%d", v)
	intMap := cache.HgetAllStringMap(key)
	log.Printf("HgetAllIntMap:%s", intMap[key])
	for k := range intMap {
		delete(intMap, k)
	}
	cache.Hset(key, "2", "this is test")
	vv := cache.Hget(key, "2") // Get bytes array

	ok := cache.IsExist(key)
	if ok {

	}

	cache.Hdel(key, "2")
	ok = cache.IsExist(key)
	if ok {

	}
	cache.Set("1", 2)
	getV := cache.Get("1")
	if getV == "2" {

	}
	getBytesV, _ := cache.GetBytes("1")
	if getV == string(getBytesV) {

	}

	log.Printf("hset value:%s", vv)
	tmpMp := cache.HgetAllStringMap(key)
	if len(tmpMp) != 0 {

	}

	l := cache.Hlen(key)

	log.Printf("len:%d", l)

	cache.Hdel(key, filed)

	ok = cache.IsExist(key)
	if ok {

	}

	//get redis client
	zaddKey := "zaddKey"
	client := cache.RedisClient()
	client.ZAdd(zaddKey, redis.Z{Score: 1, Member: "梨子"},
		redis.Z{Score: 2, Member: "苹果"},
		redis.Z{Score: 3, Member: "香蕉"})
	//client.ZRem(zaddKey, "梨子")
	score := client.ZScore(zaddKey, "香蕉").Val()
	log.Printf("score:%f", score)

	//score min and max
	//offset ,count -- 进行数据分页
	//score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
	//min 和 max 可以是 -inf 和 +inf
	members, e := client.ZRangeByScore(zaddKey, redis.ZRangeBy{
		Min: "-inf", Max: "+inf",
	}).Result()
	if e == nil {
		if len(members) != 0 {
			for _, v := range members {
				log.Printf("members:%s", v)
			}
		}
	}

	log.Printf("===========")
	{
		members, e := client.ZRevRangeByScore(zaddKey, redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Result()
		if e == nil {
			if len(members) != 0 {
				for _, v := range members {
					log.Printf("members:%s", v)
				}
			}
		}

	}

	log.Printf("+++++++++++")

	{
		members, e := client.ZRevRangeByScore(zaddKey, redis.ZRangeBy{
			Max: "3", Offset: 0, Count: 10,
		}).Result()
		if e == nil {
			if len(members) != 0 {
				for _, v := range members {
					log.Printf("members:%s", v)
				}
			}
		}
	}

	{
		_ = client.ZRemRangeByScore(zaddKey, "-inf", "+inf").String()

		f, e := client.ZScore(zaddKey, "苹果").Result()
		if e == nil {
			log.Printf("%f", f)
		} else {
			log.Printf("%s", e.Error())
		}

	}

	{

		//client.SCard()
		//client.SMembers()
		//client.SCard()

	}

	//
	intValue, _ := strconv.ParseInt("12", 10, 64)
	logger.Info("parse int", zap.Int64("intValue", intValue))
	tmp := strconv.FormatInt(12, 10)
	logger.Info("format int", zap.String("format value", tmp))

	tickerChan := time.NewTicker(2 * time.Second).C
	for true {
		select {
		case <-tickerChan:
			{
				_, e := client.Ping().Result()
				if e != nil {
					log.Printf("time:%s ping error:%s", guava.GetNowDateTime(), e.Error())
				} else {
					//log.Printf("time:%s ping:%s", guava.GetNowDateTime(), v, )

					//test protobuf
					dau := &Dau{DeviceID: "1234", ChannelName: "testchannel"}
					pbKey := "pbKey"
					cacheChannel.SetProtoBuf(regin, pbKey, dau)
					dd := cacheChannel.GetBytesLevel2(regin, pbKey)
					var dauCopy Dau
					_ = proto.Unmarshal(dd, &dauCopy)
					cacheChannel.GetProtoBufLevel2(regin, pbKey, &Dau{})
					cacheChannel.GetProtoBufLevel2(regin, pbKey, &Dau{})
					dauCacheVo := cacheChannel.GetLevel1(regin, pbKey)
					if dauCacheVo == dau {

					}
				}
			}
		}
	}
}
