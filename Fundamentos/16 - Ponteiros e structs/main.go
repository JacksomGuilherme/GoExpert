package main

import (
	"fmt"
	"math/rand"
)

type Pessoa struct {
	Nome      string
	QtdPassos int
}

func (p *Pessoa) Andar() {
	p.QtdPassos += 1
	fmt.Println(fmt.Sprintf("%v deu %d passo(s)", p.Nome, p.QtdPassos))
}

func main() {

	jack := Pessoa{
		"Jacksom",
		0,
	}

	for i := 0; i < rand.Intn(100); i++ {
		jack.Andar()
	}

	fmt.Println(fmt.Sprintf("%s andou um total de %d passos", jack.Nome, jack.QtdPassos))

}
