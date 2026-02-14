package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64 //Usando ~ faz com que o Go reconheça tipos criados com base em tipos primitivos do go
}

func Soma[T Number](m map[string]T) T {
	var total T = 0

	for _, valor := range m {
		total += valor
	}

	return total
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	salarios := map[string]int{"Jack": 5000, "Iza": 10000, "Lino": 100000}
	fmt.Println(Soma(salarios))

	salarios2 := map[string]float64{"Jack": 5000.65, "Iza": 10000.20, "Lino": 1000.00}
	fmt.Println(Soma(salarios2))

	salarios3 := map[string]MyNumber{"Jack": 5000, "Iza": 10000, "Lino": 100000}
	fmt.Println(Soma(salarios3))

	fmt.Println(Compara(10, 5))
	fmt.Println(Compara(10, 10))
}
