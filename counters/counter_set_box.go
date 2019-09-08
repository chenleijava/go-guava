package counters

import (
	"sync"
)

var onceCounterSetBox sync.Once
var counterSetBox *CounterSetBox

//counter set box
type CounterSetBox struct {
	counterSets sync.Map
	l           sync.Mutex
}

//new counter set box
func NewCounterSetBox() *CounterSetBox {
	if counterSetBox != nil {
		return counterSetBox
	}
	//init once
	onceCounterSetBox.Do(func() {
		counterSetBox = &CounterSetBox{}
	})
	return counterSetBox
}

//Get counter set box
func (c *CounterSetBox) GetCounterSetBox(name string) *CounterSet {
	value, ok := c.counterSets.Load(name)
	if ok {
		return value.(*CounterSet)
	}

	//not found -> create and store
	c.l.Lock()
	value = NewCounterSet(name)
	c.counterSets.Store(name, value) //store it
	c.l.Unlock()

	return value.(*CounterSet)
}
