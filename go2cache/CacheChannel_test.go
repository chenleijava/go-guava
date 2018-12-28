package go2cache

import (
	"github.com/chenleijava/go-guava"
	"github.com/chenleijava/go-guava/zlog"
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
	regin := "user_region"
	cache := cacheChannel.GetRedisCache("user_region")
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
					proto.Unmarshal(dd,&dauCopy)


					cacheChannel.GetProtoBufLevel2(regin,pbKey,&Dau{})
					cacheChannel.GetProtoBufLevel2(regin,pbKey,&Dau{})
					dauCacheVo := cacheChannel.GetLevel1(regin, pbKey)
					if dauCacheVo == dau {

					}



				}
			}
		}
	}
}
