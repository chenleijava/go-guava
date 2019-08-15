package list

import (
	"log"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	if element, b := l.Has(3); b {
		log.Printf("%d", element.Value)
	}

	log.Printf("before length:%d", l.Len())
	l.Range(func(v interface{}) bool {
		log.Printf("foreach data:%d", v)
		return true
	})

	//clear
	l.Clear()

	log.Printf("clear after length:%d", l.Len())
}
