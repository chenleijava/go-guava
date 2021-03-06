package list

import (
	"container/list"
	"sync"
)

type List struct {
	l    *list.List
	lock sync.RWMutex
}

//init thread safe list
func NewList() *List {
	l := new(List)
	l.l = list.New()
	return l
}

//remove all element
func (l *List) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	var next *list.Element
	for e := l.l.Front(); e != nil; e = next {
		next = e.Next()
		l.l.Remove(e)
	}
}

// Range calls f sequentially for each  value present in the List.
// If f returns false, range stops the iteration.
func (l *List) Range(f func(v interface{}) bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for e := l.l.Front(); e != nil; e = e.Next() {
		if !f(e.Value) {
			break
		}
	}
}

// Range calls f sequentially for each  value present in the List.
// If f returns false, range stops the iteration.
func (l *List) RangeElement(f func(v interface{}) bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for e := l.l.Front(); e != nil; e = e.Next() {
		if !f(e) {
			break
		}
	}
}

//get list len
func (l *List) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.l.Len()
}

//get front
func (l *List) Front() *list.Element {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.l.Front()
}

//get back
func (l *List) Back() *list.Element {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.l.Back()
}

//remove
func (l *List) RemoveLast() *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	last := l.l.Back()
	if last != nil {
		l.l.Remove(last)
		return last
	}
	return nil
}

//push element back
func (l *List) PushBack(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.l.PushBack(v)
}

//push element back
func (l *List) PushFront(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.l.PushFront(v)
}

//remove
func (l *List) Remove(v *list.Element) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.l.Remove(v)
}

//has
//return e
func (l *List) Has(v interface{}) (*list.Element, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for e := l.l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return e, true
		}
	}
	return nil, false
}
