package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Second}

	resp, err := c.Get("https://official-joke-api.appspot.com/random_joke")
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
