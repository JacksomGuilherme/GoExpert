package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	jack := Cliente{
		"Jacksom",
		26,
		true,
	}

	fmt.Println(jack)
	fmt.Println(fmt.Sprintf("Nome: %s, Idade: %d, Ativo: %v", jack.Nome, jack.Idade, jack.Ativo))

}
