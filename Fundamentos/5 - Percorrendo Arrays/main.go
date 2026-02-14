package main

import "fmt"

func main() {

	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for index, valor := range meuArray {
		fmt.Println(fmt.Sprintf("Index: %d - Valor: %d", index, valor))
	}

}
