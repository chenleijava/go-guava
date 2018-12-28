package goroutine_pool

import (
	"fmt"
	"log"
	"testing"
)

func TestAsyncGoRoutinePool(t *testing.T) {
	pool := NewAsyncPool(10, 1024)
	// release resources used by pool
	defer pool.Release()
	var j int
	//wait job count
	pool.WaitCount(10)
	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i
		//add job to queue
		pool.JobQueue <- func() {
			log.Printf("num:%d", j)
			fmt.Printf("I am worker! Number %d\n", count)
			//done one job
			pool.JobDone()
		}
	}
	pool.WaitAll()
	log.Printf("main goroutine done jobs ,hook >>>>>")
}
