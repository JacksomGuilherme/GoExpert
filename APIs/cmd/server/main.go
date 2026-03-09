package main

import (
	"fmt"

	"github.com/JacksomGuilherme/GoExpert/APIs/configs"
)

func main() {
	config := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
