package math

/*
Tudo que estiver com letra maiúscula na primeira letra será exportado,
já o que estiver minúsculo não será exportado.
*/

type math struct {
	a int
	b int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m *math) Add() int {
	return m.a + m.b
}
