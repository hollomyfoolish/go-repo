package gchan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 1)
	fmt.Println("testing ...")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go feed(ch, &wg)
	go retrive(ch, &wg)

	wg.Wait()
}

func feed(ch chan<- int, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	ch <- 1
	fmt.Println("feed done")
	wg.Done()
}

func retrive(ch <-chan int, wg *sync.WaitGroup) {
	fmt.Println("try to retrive")
	d := <-ch
	fmt.Printf("get: %d\n", d)
	wg.Done()
}
