package go2cache

import (
	"fmt"
	"github.com/chenleijava/go-guava"
	"github.com/magiconair/properties"
	"log"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once
var cfg *RedisConfig

//load go2cache config
func GetConfig() *RedisConfig {
	if cfg == nil {
		once.Do(func() {
			//loading redis config
			pp, ee := properties.LoadFile(guava.ExePath()+"/resources/config/go2cache.properties", properties.UTF8)
			if ee != nil {
				//load form base config
				pp, ee = properties.LoadFile(guava.ExePath()+"/config/go2cache.properties", properties.UTF8)
				if ee != nil {
					log.Panicf("load redis config bad error:%s , config must be in `/resources` or `/config` path !", ee.Error())
				}
			}
			var _cfg RedisConfig
			if err := pp.Decode(&_cfg); err != nil {
				log.Fatalf("load config  error:%s", err.Error())
			}
			pp.ClearComments()

			go2cacheRegions := _cfg.Go2cacheRegions
			if go2cacheRegions != "" {
				tmp := strings.Split(go2cacheRegions, ";")
				for _, v := range tmp {
					if v == "" {
						continue
					}
					lv := strings.Split(v, ",")
					s, _ := strconv.Atoi(lv[1])
					_cfg.Regions = append(_cfg.Regions, &Region{Name: lv[0], Size: s});
				}
			}
			cfg = &_cfg
		})
	}
	return cfg
}

//redis config
type RedisConfig struct {
	Addr            string    `properties:"addr"`
	Password        string    `properties:"password"`
	MinIdleConns    int       `properties:"minIdleConns"`
	PoolSize        int       `properties:"poolSize"`
	Db              int       `properties:"db"`
	Channel         string    `properties:"channel"`
	RedisNameSpace  string    `properties:"redisNameSpace"`
	Psb             bool      `properties:"psb"`
	Go2cacheRegions string    `properties:"go2CacheRegions"`
	Regions         []*Region `properties:"-"`
}

type Region struct {
	Name string //当前region名称
	Size int    //当前区域下 对应cache的容积大小
}

func (r *Region) String() string {
	return fmt.Sprintf("region:%s@%d", r.Name, r.Size)
}
