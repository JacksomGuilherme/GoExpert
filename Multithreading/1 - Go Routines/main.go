package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("%d: Task %s is running", i, name))
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(fmt.Sprintf("%d: Task %s is running", i, "anonymous"))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
