package main

import "fmt"

func main() {
	fmt.Println("primeiro")
	defer fmt.Println("segundo")
	fmt.Println("terceiro")
}
