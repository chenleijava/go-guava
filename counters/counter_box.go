package counters

import "sync"

//once
var onceCounterBox sync.Once
var counterBox *CounterBox

//cache all counter
type CounterBox struct {
	counters sync.Map
	l        sync.Mutex
}

//new counter box
func NewCounterBox() *CounterBox {
	if counterBox != nil {
		return counterBox
	}
	//init counter box
	onceCounterBox.Do(func() {
		counterBox = &CounterBox{}
	})
	return counterBox
}

// all counters
// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (c *CounterBox) Range(f func(name, value interface{}) bool) {
	c.counters.Range(f)
}

// GetCounter returns a counter of given name,
// if doesn't exist than create.
func (c *CounterBox) GetCounter(name string) *Counter {
	value, ok := c.counters.Load(name)
	if ok {
		return value.(*Counter)
	}

	//not found -> create and store
	c.l.Lock()
	value = NewCounter(name)
	c.counters.Store(name, value)
	c.l.Unlock()

	return value.(*Counter)
}
