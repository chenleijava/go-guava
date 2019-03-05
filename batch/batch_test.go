package batch

import (
	"go.uber.org/atomic"
	"log"
	"sync"
	"testing"
	"time"
)



func TestNewBatch2(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	// Create a batch with a given size & duration. If the number of items hits the
	// configured maximum or the given timeout is exceeded, the items are written
	// to a channel.
	batch := New(2, time.Second*1, func(values []interface{}) {
		// do something
		for _, v := range values {
			log.Printf("%s", v)
		}
	})
	var c atomic.Int32
	for true {
		// Add some items to the batch
		batch.Add("test")
		batch.Add("测试数据")
		batch.Add("测试数据2")
		d := c.Inc()
		if d%2 == 0 {
			 batch.ForceFlushChan() //
			log.Printf("ForceFlush>>>")
		}
		time.Sleep(time.Second * 2)
	}
	wait.Wait()

}
