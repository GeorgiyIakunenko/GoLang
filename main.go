package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//fmt.Printf("concurrency")

	done := make(chan bool)
	fmt.Printf("type of channel : %T\n", done)

	go hello(done)

	<-done

	runtime.GOMAXPROCS(1)
	go numbers()
	go letters()

	fmt.Printf("main goruitine")
}

func numbers() {
	for i := 1; i < 20; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d ", i)
	}
}

func letters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c ", i)
	}
}

func hello(done chan bool) {
	fmt.Printf("hello world goroutine\n")
	for i := 1; i < 4; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	done <- true
}
