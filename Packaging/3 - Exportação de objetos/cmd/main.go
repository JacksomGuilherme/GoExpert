package main

import (
	"fmt"
	"pacotes/math"
)

func main() {
	m := math.NewMath(1, 5)
	fmt.Println(m.Add())
}
