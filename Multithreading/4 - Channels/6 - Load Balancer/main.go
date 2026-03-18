package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	qtdWorkers := 100
	for i := 0; i < qtdWorkers; i++ {
		go worker(i+1, data)
	}

	for i := 0; i < 1000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Println(fmt.Sprintf("Worker %d received %d", workerId, x))
		time.Sleep(time.Second)
	}
}
