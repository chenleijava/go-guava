package batch

import (
	"sync"
	"time"
)

type (
	// The Batch type contains items added to it and writes to a channel when the batch is full or
	// has been runing a configured time.
	Batch struct {
		maxSize int
		maxWait time.Duration

		inner []interface{}

		notify         chan []interface{}
		add            chan interface{}
		close          chan bool
		forceFlushChan chan bool
		f              func([]interface{}) //logic hand
		mu             sync.Mutex          // lock max wait time
	}
)

// New creates a new Batch with a given size & run time. The batch writes to a channel when
// either the batch is full or the batch has runed the configured time.
func New(maxSize int, maxWait time.Duration, f func(values []interface{})) *Batch {
	batch := &Batch{
		maxSize:        maxSize,
		maxWait:        maxWait,
		inner:          make([]interface{}, 0, maxSize),
		notify:         make(chan []interface{}),
		add:            make(chan interface{}),
		close:          make(chan bool),
		forceFlushChan: make(chan bool),
		f:              f,
	}

	//loop and handle
	//add ,forceFlush,close
	go batch.run()

	//ticker flush data
	go batch.ticker()

	//set flush
	batch.setFlush()

	return batch
}

//block and run data to handle
func (b *Batch) setFlush() {
	go func() {
		for items := range b.notify {
			b.f(items)
			items = nil //set nil
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

//ticker to flush data
func (b *Batch) ticker() {
	//flushData time chan
	flushDataTicker := time.Tick(b.maxWait)
	for {
		select {
		// If we've reached the maximum run time
		// Write batch contents to channel,
		// If  len(b.inner) >0, write batch
		// contents to channel, clear batched item and add new
		// item to empty batch.
		case <-flushDataTicker:
			b.mu.Lock()
			if len(b.inner) > 0 {
				b.flush()
			}
			b.mu.Unlock()
		}
	}
}

// Close causes the batch to stop checking its size & duration. Should be used when the
// batch is no longer required.
func (b *Batch) Close() {
	b.close <- true
}

//
func (b *Batch) run() {
	for {
		select {
		// If an item has been added to the batch.
		// If we've reached the maximum batch size, write batch
		// contents to channel, clear batched item and add new
		// item to empty batch.
		case item := <-b.add:
			b.mu.Lock()
			b.inner = append(b.inner, item)
			if len(b.inner) == b.maxSize {
				b.flush()
			}
			b.mu.Unlock()
		// If we've force flush chan and len(b.inner) !=0, write batch
		// contents to channel, clear batched item and add new
		// item to empty batch.
		case <-b.forceFlushChan:
			b.mu.Lock()
			if len(b.inner) > 0 {
				b.flush()
			}
			b.mu.Unlock()
		//if close ,data flush
		// If the batch has been closed, wipe the batch clean,
		// close channels & exit the loop.
		case <-b.close:
			b.mu.Lock()
			if len(b.inner) > 0 {
				b.flush()
			}
			close(b.notify)
			close(b.close)
			close(b.add)
			close(b.forceFlushChan)
			b.maxSize = 0
			b.maxWait = 0
			b.inner = nil
			b.mu.Unlock()
			return
		}
	}
}

//flush data to chan
//thread safe
func (b *Batch) flush() {
	b.notify <- b.inner   // flush to notify chain ,pass copy value
	b.inner = b.inner[:0] //reset
	//log.Printf("reset done  len:%d cap:%d",len(b.inner),cap(b.inner))
}
