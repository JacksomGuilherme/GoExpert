package main

import "fmt"

func main() {

	total := func() int {
		return soma(2, 5, 3, 4) * 2
	}()

	fmt.Println(fmt.Sprintf("O total é %d", total))

}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
