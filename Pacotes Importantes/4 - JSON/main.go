package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Usuario struct {
	Nick  string `json:"n"` //se colocar um - o json ignora o campo
	Email string `json:"e"`
}

func main() {
	/* Da forma a seguir ao fazer o Marshal é possível guardar o valor em json */
	usuario := Usuario{
		"Jack",
		"jack@golang.com",
	}

	res, err := json.Marshal(usuario)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	/* ------------------------------------ */

	/* Da forma a seguir ao fazer Encode o encoder apenas envia o valor para o que foi passado para ele */
	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(usuario)
	if err != nil {
		panic(err)
	}

	/* ------------------------------------ */

	/* O codigo a seguir transofrma um array de bytes correspondentes ao json em um strcut */
	jsonPuro := []byte(`{"Nick": "lino", "Email": "lino@golang.com"}`)
	var user Usuario
	json.Unmarshal(jsonPuro, &user)

	fmt.Println(user)

	/* ------------------------------------ */

	/* O codigo a seguir usa os "alias" incluidos no struct para conseguir concluir os binds dos valores ao fazer Unmarshal */
	jsonPuro = []byte(`{"n": "izzy", "e": "izzy@golang.com"}`)
	var user2 Usuario
	json.Unmarshal(jsonPuro, &user2)

	fmt.Println(user2)

	/* ------------------------------------ */
}
