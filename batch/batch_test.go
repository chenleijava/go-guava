package batch

import (
	"github.com/herryg91/gobatch"
	"github.com/robfig/cron"
	"log"
	"sync"
	"testing"
	"time"
)

/*
select {
case v1 := <-c1:
fmt.Printf("received %v from c1\n", v1)
case v2 := <-c2:
fmt.Printf("received %v from c2\n", v1)
case c3 <- 23:
fmt.Printf("sent %v to c3\n", 23)
default:
fmt.Printf("no one was ready to communicate\n")
}
上面这段代码中，select 语句有四个 case 子语句，前两个是 receive 操作，第三个是 send 操作，
最后一个是默认操作。代码执行到 select 时，case 语句会按照源代码的顺序被评估，且只评估一次，评估的结果会出现下面这几种情况：
1.除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
2.除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
3.如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
4.如果没有 default，那么 代码块会被阻塞，直到有一个 case 通过评估；否则一直阻塞
*/

func TestNewBatch2(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	// Create a batch with a given size & duration. If the number of items hits the
	// configured maximum or the given timeout is exceeded, the items are written
	// to a channel.
	batch := New(100, time.Second*10, func(values []interface{}) {
		// do something
		for _, v := range values {
			log.Printf("%s", v)
		}
		log.Printf("%d", len(values))
	})
	//batch.Add("close data to show")
	//time.Sleep(time.Second*1)
	//batch.Close()

	//添加的频度大于了定时任务
	c:=cron.New()
	_ = c.AddFunc("*/1 * * * * ?", func() {
		batch.Add("测试数据2")
	})
	c.Start()

	wait.Wait()
}

func TestNew(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	batch:=gobatch.NewMemoryBatch(func(workerID int, datas []interface{}) (err error) {
		for _, v := range datas {
			log.Printf("workID:%d data:%s",workerID ,v)
		}
		return nil
	},100,time.Second*1,1)
	log.Printf("begin time>>>>>")
	for true {
		// Add some items to the batch
		_ = batch.Insert("测试数据2")
		time.Sleep(time.Second * 5)
	}
	wait.Wait()
}
