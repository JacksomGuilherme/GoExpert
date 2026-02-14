package main

import (
	"fmt"
	"pacotes_e_modulos/matematica"
)

func main() {

	s := matematica.Soma(10, 20)
	fmt.Println(fmt.Sprintf("O resultado da soma é: %v", s))

}
