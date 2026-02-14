package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool

	Endereco // Dessa forma é composição
	//Address Endereco -> Dessa forma é apenas um atributo de Cliente do tipo Endereco
}

func main() {

	jack := Cliente{
		"Jacksom",
		26,
		true,
		Endereco{
			"Rua dos Bobos",
			0,
			"São Paulo",
			"SP",
		},
	}

	fmt.Println(jack)
	fmt.Println(fmt.Sprintf("Nome: %s, Idade: %d, Ativo: %v", jack.Nome, jack.Idade, jack.Ativo))

}
