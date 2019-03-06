package batch

import (
	"time"
)

type (
	// The Batch type contains items added to it and writes to a channel when the batch is full or
	// has been waiting a configured time.
	Batch struct {
		maxSize int
		maxWait time.Duration

		empty []interface{}
		inner []interface{}

		notify         chan []interface{}
		add            chan interface{}
		close          chan bool
		forceFlushChan chan bool
		f              func([]interface{}) //logic hand
	}
)

// New creates a new Batch with a given size & wait time. The batch writes to a channel when
// either the batch is full or the batch has waited the configured time.
func New(maxSize int, maxWait time.Duration, f func(values []interface{})) *Batch {
	batch := &Batch{
		maxSize:        maxSize,
		maxWait:        maxWait,
		inner:          make([]interface{}, 0, maxSize),
		empty:          make([]interface{}, 0, maxSize),
		notify:         make(chan []interface{}),
		add:            make(chan interface{}),
		close:          make(chan bool),
		forceFlushChan: make(chan bool),
		f:              f,
	}

	//loop and handle
	go batch.wait()

	//set flush
	batch.setFlush()

	return batch
}

//block and wait data to handle
func (b *Batch) setFlush() {
	go func() {
		for items := range b.notify {
			b.f(items)
		}
	}()
}

func (b *Batch) ForceFlushChan() {
	b.forceFlushChan <- true
}

// Add adds an item to the batch.
func (b *Batch) Add(item interface{}) {
	b.add <- item
}

// Close causes the batch to stop checking its size & duration. Should be used when the
// batch is no longer required.
func (b *Batch) Close() {
	b.close <- true
}

//
func (b *Batch) wait() {
	for {
		select {
		// If we've reached the maximum wait time
		case <-time.Tick(b.maxWait):
			// Write batch contents to channel,
			// If  len(b.inner) !=0, write batch
			// contents to channel, clear batched item and add new
			// item to empty batch.
			if len(b.inner) > 0 {
				b.flush()
			}
		// If an item has been added to the batch.
		case item := <-b.add:
			b.inner = append(b.inner, item)
			// If we've reached the maximum batch size, write batch
			// contents to channel, clear batched item and add new
			// item to empty batch.
			if len(b.inner) == b.maxSize {
				b.flush()
			}
		case <-b.forceFlushChan:
			// If we've force flush chan and len(b.inner) !=0, write batch
			// contents to channel, clear batched item and add new
			// item to empty batch.
			if len(b.inner) > 0 {
				b.flush()
			}
		case <-b.close:
			//if close ,data flush
			if len(b.inner) > 0 {
				b.flush()
			}

			// If the batch has been closed, wipe the batch clean,
			// close channels & exit the loop.
			close(b.notify)
			close(b.close)
			close(b.add)
			close(b.forceFlushChan)
			b.maxSize = 0
			b.maxWait = 0
			b.inner = nil
			b.empty = nil
			return
		}
	}
}

//flush data to chan
func (b *Batch) flush() {
	b.notify <- b.inner // flush to notify chain
	b.inner = b.empty   //reset
}
