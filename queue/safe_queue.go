package queue

import (
	"github.com/eapache/queue"
	"sync"
)

//thread safe queue
type SafeQueue struct {
	Queue *queue.Queue
	l     sync.RWMutex
}

func NewQueue() *SafeQueue {
	return &SafeQueue{
		Queue: queue.New(),
	}
}

func (s *SafeQueue) Add(e interface{}) {
	s.l.Lock()
	s.Queue.Add(e)
	s.l.Unlock()
}

func (s *SafeQueue) Pop() interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.Queue.Remove()
}

func (s *SafeQueue) Peek() interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	return s.Queue.Peek()
}

func (s *SafeQueue) Length() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.Queue.Length()
}
