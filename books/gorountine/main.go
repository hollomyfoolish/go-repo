package main

import(
	"fmt"
	"time"
	"context"
)

type Foo struct{
	Name string
	Age int
}

type Too struct{
	Foo
	Type string
}

type Yoo struct{
	Too
	Signal string
}

func (f Foo) GetDisplayName() string {
	return fmt.Sprintf("Name: %s, Age: %d", f.Name, f.Age)
}

func (t Too) ToString() string {
	return t.Name
}

func main () {
	ch := make(chan int,1)

	go check(ch)
 
	val := <- ch

	fmt.Printf("No.1 value: %v\n", val)
	// val = <- ch
	// fmt.Printf("No.2 value: %v\n", val)

	fmt.Printf("%v\n", context.Background())

	f := Foo{
		Name: "Foo",
		Age: 18,
	}

	fmt.Printf("%v\n", f.GetDisplayName())

	t := Too{
		Foo: f,
		Type: "T",
	}
	fmt.Printf("%v\n", t.ToString())
	fmt.Printf("%v\n", t.Foo.Name)
	fmt.Printf("%v\n", t.Name)

	y := Yoo{
		Too: t,
		Signal: "s",
	}
	fmt.Printf("%v\n", y.Name)

	checkJob()
}

func checkJob(){
	timeout := time.After(1500 * time.Millisecond)
	ch := make(chan bool)

	select{
	case t := <- timeout:
		fmt.Printf("timeout: %v\n", t)
	case t:= <- ch:
		fmt.Printf("job done: %v\n", t)
	}
	fmt.Println("job check done")

	chi := make(chan Foo)
	go closeChan(chi)
	v := <- chi
	fmt.Println(v)
	close(chi)
	v = <- chi
	fmt.Println(v)
}



func closeChan(ch chan <- Foo){
	time.Sleep(1 * time.Second)
	ch <- Foo{
		Name: "Foo",
	}
	// close(ch)
}

func doJob(ch chan<- bool){
	time.Sleep(1 * time.Second)
	ch <- true
}

func check(ch chan<- int){
	time.Sleep(1 * time.Second)
	ch <- 2
}