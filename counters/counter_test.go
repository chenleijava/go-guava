package counters

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCounter_Incr(t *testing.T) {
	c := NewCounterBox()
	var counter = c.GetCounter("abc123")
	counter.Incr()
	v := counter.IncrBy(2)
	counter.Decr()
	v = c.GetCounter("abc123").DecrBy(1)
	if v == 1 {
		log.Printf("%d", v)
	}
}

func TestCounter_Set(t *testing.T) {
	counterSet := NewCounterSet("1")
	counterSet.Add(1)
	assert.Equal(t, true, counterSet.Contains(1))
	assert.Equal(t, 1, counterSet.Len())
}

func BenchmarkCountersCached(b *testing.B) {
	b.StopTimer()
	e := make(chan bool)
	c := NewCounterBox()
	f := func(b *testing.B, c *CounterBox, e chan bool) {
		x := c.GetCounter("abc123")
		y := c.GetCounter("def456")
		z := c.GetCounter("ghi789")
		for i := 0; i < b.N; i++ {
			x.IncrBy(5)
			y.IncrBy(5)
			z.IncrBy(5)
			x.IncrBy(5)
			y.IncrBy(5)
			z.IncrBy(5)
		}
		e <- true
	}
	b.StartTimer()
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)

	<-e
	<-e
	<-e
	<-e
	<-e
	<-e
	<-e
}

func BenchmarkNewCounterBox(b *testing.B) {
	b.StopTimer()
	e := make(chan bool)
	c := NewCounterBox()
	NewCounterBox()
	f := func(b *testing.B, c *CounterBox, e chan bool) {
		for i := 0; i < b.N; i++ {
			c.GetCounter("abc123").IncrBy(5)
			c.GetCounter("def456").IncrBy(5)
			c.GetCounter("ghi789").IncrBy(5)
			c.GetCounter("abc123").IncrBy(5)
			c.GetCounter("def456").IncrBy(5)
			c.GetCounter("ghi789").IncrBy(5)
		}
		e <- true
	}
	b.StartTimer()
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)
	go f(b, c, e)

	<-e
	<-e
	<-e
	<-e
	<-e
}

func TestNewCounterSetBox(b *testing.T) {
	NewCounterSetBox().GetCounterSetBox("x").Add(1)
	NewCounterSetBox().GetCounterSetBox("x").Remove(1)
}
