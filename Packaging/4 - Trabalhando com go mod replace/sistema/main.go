package main

import (
	"fmt"

	"github.com/jacksomguilherme/goexpert/pacote-math/math"
)

func main() {
	m := math.NewMath(1, 5)
	fmt.Println(m.Add())
}
