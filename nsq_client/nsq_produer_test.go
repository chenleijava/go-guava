package nsq_client

import (
	"sync"
	"testing"
	"time"
)

const testTopic = "test"

func TestInitProducer(t *testing.T) {
	strIP1 := "127.0.0.1:4150"
	p := InitProducer(strIP1, testTopic)
	for true {
		time.Sleep(time.Second)
		command := []byte("123")
		err := p.Publish(command)
		if err != nil {
			//retry send message
		}
	}
}

//https://leetcode-cn.com/problems/print-in-order/
//wait test close
var wt sync.WaitGroup
var m sync.Mutex
var cc = sync.NewCond(&m)
var flg int32

func first() {
	cc.L.Lock()
	println("first")
	flg = 2
	cc.Broadcast()
	cc.L.Unlock()
}

func second() {
	cc.L.Lock()
	for flg != 2 {
		//println("wait flg == 2 ")
		cc.Wait()
	}
	println("second")
	flg = 3
	cc.Broadcast()
	wt.Done()
	cc.L.Unlock()
}

func third() {
	cc.L.Lock()
	for flg != 3 {
		//println("wait flg ==3")
		cc.Wait()
	}
	println("third")
	wt.Done()
	cc.L.Unlock()
}

//
func TestInitProducer2(t *testing.T) {
	println(1234 ^ 5678)
	//wt.Add(2)
	//go first()
	//go third()
	//go second()
	//wt.Wait()

}
