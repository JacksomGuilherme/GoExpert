package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stdout, "Erro fazer a requisição %v\n", err)
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(os.Stdout, "Erro ao ler a resposta %v\n", err)
		}
		defer req.Body.Close()

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprint(os.Stdout, "Erro ao converter o json %v\n", err)
		}

		fmt.Println(data)

	}

}
