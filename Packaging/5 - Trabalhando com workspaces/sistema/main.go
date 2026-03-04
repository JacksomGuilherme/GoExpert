package main

import (
	"fmt"

	math "github.com/jacksomguilherme/goexpert/pacote-math"
)

func main() {
	m := math.NewMath(1, 5)
	fmt.Println(m.Add())
}
