package counter

import (
	mapset "github.com/deckarep/golang-set"
	"sync/atomic"
)

//base counter
type Counter struct {
	count int64
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

//new counter
func NewCounter() *Counter {
	return &Counter{count: 0}
}

//base counter with name
type CounterModule struct {
	Name    string
	Counter *Counter
}

//new counter module
func NewCounterModule(name string) *CounterModule {
	return &CounterModule{
		Name:    name,
		Counter: NewCounter(),
	}
}

//counter set
type CounterSetI interface {
	// Adds an element to the set. Returns whether
	// the item was added.
	Add(i interface{}) bool
	// Returns the number of elements in the set.
	Len() int
	// Returns whether the given items
	// are all in the set.
	Contains(i ...interface{}) bool
}

//base counter set
type CounterSet struct {
	set  mapset.Set
	Name string
}

//new counter set
func NewCounterSet(name string) *CounterSet {
	return &CounterSet{
		set:  mapset.NewSet(),
		Name: name,
	}
}

// Adds an element to the set. Returns whether
// the item was added.
func (c *CounterSet) Add(i interface{}) bool {
	return c.set.Add(i)
}

// Returns the number of elements in the set.
func (c *CounterSet) Len() int {
	return c.set.Cardinality()
}

// Returns whether the given items
// are all in the set.
func (c *CounterSet) Contains(i ...interface{}) bool {
	return c.set.Contains(i...)
}
