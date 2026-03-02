package main

import (
	"fmt"
	"pacotes/math"
)

func main() {
	m := math.Math{A: 1, B: 5}
	fmt.Println(m.Add())
}
