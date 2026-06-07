package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var p fastjson.Parser
	jsonData := `{ "user": {"name": "Lino", "age": 1} }`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJSON := v.Get("user").String()

	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("name=%s", user.Name))
	fmt.Println(fmt.Sprintf("age=%v", user.Age))

}
