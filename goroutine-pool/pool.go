package goroutine_pool

import (
	"github.com/Jeffail/tunny"
	"github.com/ivpusic/grpool"
)

//pool := NewAsyncPool(10, 1024)
//// release resources used by pool
//defer pool.Release()
//var j int
////wait job count
//pool.WaitCount(10)
//// submit one or more jobs to pool
//for i := 0; i < 10; i++ {
//count := i
////add job to queue
//pool.JobQueue <- func() {
//log.Printf("num:%d", j)
//fmt.Printf("I am worker! Number %d\n", count)
////done one job
//pool.JobDone()
//}
//}
//pool.WaitAll()
//log.Printf("main goroutine done jobs ,hook >>>>>")
//create async goroutine pool
func NewAsyncPool(routineSize, jobSize int) *grpool.Pool {
	return grpool.NewPool(routineSize, jobSize)
}

//https://github.com/Jeffail/tunny
//Tunny is a Golang library for spawning and managing a goroutine pool, allowing you to limit work
//coming from any number of goroutines with a synchronous API.
//A fixed goroutine pool is helpful when you have work coming from an arbitrary number of asynchronous sources,
//but a limited capacity for parallel processing. For example, when processing jobs from HTTP requests
//that are CPU heavy you can create a pool with a size that matches your CPU count.
func NewSyncPool(routineSize int, process func(payload interface{}) interface{}) *tunny.Pool {
	return tunny.NewFunc(routineSize, process)
}
