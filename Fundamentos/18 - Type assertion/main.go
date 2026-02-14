package main

import "fmt"

func main() {

	var minhaVar interface{} = "Hello, World!"
	fmt.Println(minhaVar.(string))

	res, ok := minhaVar.(int)

	fmt.Println(fmt.Sprintf("O valor de res é %v e o valor de ok é %v", res, ok))

}
