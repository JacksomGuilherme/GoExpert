package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(soma(5, 22))

	valor, erro := soma2(56, 22)

	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(valor)
}

func soma(a, b int) int {
	return a + b
}

func soma2(a, b int) (int, error) {
	if (a + b) > 50 {
		return 0, errors.New("A soma é maior do que 50")
	}
	return a + b, nil
}
