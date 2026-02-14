package main

import "fmt"

func main() {

	salarios := map[string]int{"jack": 1000, "iza": 4000, "lino": 50000}
	// sal := make(map[string]int)

	// sal1 := map[string]int{}

	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Println(fmt.Sprintf("Colaborador: %s, Salário: %d", nome, salario))
	}

	fmt.Printf("\n")

	delete(salarios, "jack")
	for nome, salario := range salarios {
		fmt.Println(fmt.Sprintf("Colaborador: %s, Salário: %d", nome, salario))
	}

	fmt.Printf("\n")

	salarios["jack"] = 2000
	for nome, salario := range salarios {
		fmt.Println(fmt.Sprintf("Colaborador: %s, Salário: %d", nome, salario))
	}

	fmt.Printf("\n")

	for _, salario := range salarios {
		fmt.Println(fmt.Sprintf("Salário: %d", salario))
	}

}
