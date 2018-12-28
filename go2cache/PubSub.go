/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package go2cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

//
type PubSub struct {
	Client       *redis.Client
	Channel      string
	CacheChannel *CacheChannel
	Region       string
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
					log.Printf("evict key :%s region:%s", cmd.Keys, cmd.Region)
					p.CacheChannel.Evict(cmd.Region, cmd.Keys)
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
