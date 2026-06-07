package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser
	jsonData := `{"foo": "bar", "num": 123, "bool": true, "arr": [1, 2, 3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
	fmt.Println(fmt.Sprintf("foo=%s", v.GetStringBytes("foo")))
	fmt.Println(fmt.Sprintf("num=%d", v.GetInt("num")))
	fmt.Println(fmt.Sprintf("bool=%v", v.GetBool("bool")))
	fmt.Println(fmt.Sprintf("arr=%v", v.GetArray("arr")))

	a := v.GetArray("arr")
	for i, value := range a {
		fmt.Println(fmt.Sprintf("Index: %d, Value: %v", i, value))
	}

}
