package main

import "fmt"

func main() {

	a := 1
	b := 2
	c := 3

	/* operadores de comparação:
	- >
	- <
	- >=
	- <=
	- !=
	- ==
	*/
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}

	/* operadores lógicos
	- &&
	- ||
	*/
	if a > b && c > a {
		fmt.Println("a > b && c > a")
	}

	if a > b || c > a {
		fmt.Println("a > b || c > a")
	}

	switch a {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	default:
		fmt.Println("d")
	}

}
