package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"type":"general","setup":"What do you call a thieving alligator?","punchline":"A crookodile!","id":451}`))

	resp, err := c.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
