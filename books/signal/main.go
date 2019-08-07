package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGTERM)
	signal.Notify(s, syscall.SIGINT)
	signal.Notify(s, syscall.SIGKILL)

	go gracefulExt(s)

	pid := os.Getpid()
	for {
		fmt.Printf("%d is running ...\n", pid)
		time.Sleep(1 * time.Second)
	}
}

func gracefulExt(s <-chan os.Signal) {
	ss := <-s
	fmt.Printf("received signal %v, waiting for exist\n", ss)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
