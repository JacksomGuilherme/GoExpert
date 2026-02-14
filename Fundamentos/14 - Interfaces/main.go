package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println(fmt.Sprintf("Cliente %s desativado", c.Nome))
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

	Desativacao(jack)
	fmt.Println(fmt.Sprintf("Nome: %s, Idade: %d, Ativo: %v", jack.Nome, jack.Idade, jack.Ativo))

}
