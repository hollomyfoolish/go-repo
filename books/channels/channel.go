package main

import(
	"fmt"
	"time"
	"reflect"
)

type MyInt int

func main(){
	ch := make(chan int, 5)
	signal := make(chan bool)

	go feed(ch)
	go fetch(ch, signal)
	go foo(ch)

	<- signal

	mi := new(MyInt)
	fmt.Printf("feed: %d\n", mi)
}

func foo(ch chan int){
	fmt.Printf("channel type: %v\n", reflect.TypeOf(ch).Name())
}

func feed(ch chan<- int){
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("feed: %d\n", i)
		ch<- i
	}
	fmt.Printf("close channel\n")
	close(ch)
}

func fetch(ch <-chan int, signal chan<- bool){
	for ; ; {
		time.Sleep(2 * time.Second)
		i := <- ch
		fmt.Printf("fetch: %d\n", i)
		if i == 0 {
			break
		}
	}
	signal<- true
}