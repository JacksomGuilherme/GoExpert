package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("%d: Task %s is running", i, name))
		time.Sleep(time.Second)
		waitGroup.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(fmt.Sprintf("%d: Task %s is running", i, "anonymous"))
			time.Sleep(time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
