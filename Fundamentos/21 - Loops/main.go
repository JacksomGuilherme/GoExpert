package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=========================================")

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		fmt.Println(k, v)
	}

	fmt.Println("=========================================")

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=========================================")

	//for vazio faz o mesmo que um while(true)
	// for {
	// 	fmt.Println("Hello, World!")
	// }
}
