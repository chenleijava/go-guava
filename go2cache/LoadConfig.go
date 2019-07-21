package go2cache

import (
	"fmt"
	"github.com/chenleijava/go-guava"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
			configData, ee := ioutil.ReadFile(guava.ExePath()+"/resources/config/go2cache.yaml")
			if ee != nil {
				//load form base config
				tmp,ee:=ioutil.ReadFile(guava.ExePath()+"/config/go2cache.yaml")
				if ee != nil {
					log.Panicf("load redis config bad error:%s , config must be in `/resources` or `/config` path !", ee.Error())
				}
				configData=tmp
			}

			cfg=&RedisConfig{} //init config object

			err := yaml.Unmarshal(configData, cfg)
			if err!=nil{
				log.Fatalf("load redis config bad err:%s",err.Error())
			}

			go2cacheRegions := cfg.Go2CacheRegions
			if go2cacheRegions != "" {
				tmp := strings.Split(go2cacheRegions, ";")
				for _, v := range tmp {
					if v == "" {
						continue
					}
					lv := strings.Split(v, ",")
					s, _ := strconv.Atoi(lv[1])
					cfg.Regions = append(cfg.Regions, &Region{Name: lv[0], Size: s});
				}
			}
		})
	}
	return cfg
}


type RedisConfig struct {
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
	MinIdleConns int `yaml:"minIdleConns"`
	PoolSize int `yaml:"poolSize"`
	Db int `yaml:"db"`
	Channel string `yaml:"channel"`
	RedisNameSpace string `yaml:"redisNameSpace"`
	Psb bool `yaml:"psb"`
	Go2CacheRegions string `yaml:"go2CacheRegions"`
	Regions         []*Region `yaml:"-"`
}

type Region struct {
	Name string //当前region名称
	Size int    //当前区域下 对应cache的容积大小
}

func (r *Region) String() string {
	return fmt.Sprintf("region:%s@%d", r.Name, r.Size)
}
