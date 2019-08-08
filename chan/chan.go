package gchan

import (
	"fmt"
	"sync"
	"time"
)

func ExecuteJobs(jobs []func(), maxThreads int) {
	qSize := len(jobs)
	if qSize == 0 {
		fmt.Println("no job to execute")
		return
	}
	ch := make(chan func(), qSize)
	wg := sync.WaitGroup{}
	wg.Add(qSize)
	for i := 0; i < maxThreads; i++ {
		go execute(&wg, ch)
	}

	for _, job := range jobs {
		ch <- job
	}
	wg.Wait()
	close(ch)
}

func execute(wg *sync.WaitGroup, q <-chan func()) {
	for {
		job, ok := <-q
		if ok {
			job()
			wg.Done()
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
}
