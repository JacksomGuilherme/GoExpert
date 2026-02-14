package matematica

//com S maiúsculo o método se torna publico
func Soma[T int | float64](a, b T) T {
	return a + b
}

//com S minúsculo o método se torna privado
func subtrair[T int | float64](a, b T) T {
	return a + b
}

//a mesma regra vale para strcuts e types criados, com letra maiúscula se torna public e minuúscula se torna privado
type A int
type b int

type number interface {
	int | float64
}
