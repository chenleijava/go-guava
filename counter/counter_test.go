package counter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var counter = NewCounter()

func TestCounter_Incr(t *testing.T) {
	counter.Incr()
	v := counter.IncrBy(2)
	counter.Decr()
	v = counter.DecrBy(1)
	if v == 1 {
		log.Printf("%d", v)
	}
}

func TestCounter_Set(t *testing.T) {
	counterSet := NewCounterSet("test")
	counterSet.Add(1)
	assert.Equal(t,true,counterSet.Contains(1))
	assert.Equal(t, 1, counterSet.Len())

}
