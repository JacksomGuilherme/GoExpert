package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Println(fmt.Sprintf("O tipo da variável é %T e o valor é %v", t, t))
}
