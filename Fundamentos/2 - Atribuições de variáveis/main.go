package main

var (
	b bool = true
	c int
	d string
	e float64
)

func main() {
	var a = "Hello World" // mutável

	const x = 10 // imutável

	y := 20 // := cria, infere o tipo e inicia a variável

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	println(x)
	println(y)
}
