package main

import "fmt"

func main() {

	// Memoria -> Endereço -> Valor

	a := 10

	fmt.Println(&a) //Espaço na memória

	var ponteiro *int = &a
	fmt.Println(ponteiro) //Ponteiro pro endereço de memória

	fmt.Println(a) //Valor

	*ponteiro = 20 //Muda o valor pelo endereço de memória
	fmt.Println(a)

	//Outra forma de criar um ponteiro
	b := &a
	fmt.Println(b)
	*b = 30
	fmt.Println(a)
}
