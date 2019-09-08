package counters

import (
	"sync/atomic"
)

//base counter
type Counter struct {
	count int64
	name  string
}

//incr
func (c *Counter) Incr() int64 {
	return atomic.AddInt64(&c.count, 1)
}

//incr by
func (c *Counter) IncrBy(num int64) int64 {
	return atomic.AddInt64(&c.count, num)
}

//dec
func (c *Counter) Decr() int64 {
	return atomic.AddInt64(&c.count, -1)
}

//dec
func (c *Counter) DecrBy(num int64) int64 {
	return atomic.AddInt64(&c.count, -num)
}

//set
func (c *Counter) Set(num int64) {
	atomic.StoreInt64(&c.count, num)
}

//get
func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

//new counter
func NewCounter(name string) *Counter {
	return &Counter{count: 0, name: name}
}
