package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello, World!\n")
	if err != nil {
		panic(err)
	}
	tamanho, err = f.Write([]byte("Teste escrevendo no arquivo com array de bytes"))

	fmt.Println(fmt.Sprintf("Arquivo criado com sucesso! tamanho do arquivo: %d bytes", tamanho))

	defer f.Close()

	//leitura do arquivo

	/* 	Dessa foram o go já abre o arquivo e lê inteiro de uma vez
		porém existe o problema de estouro de memória caso o tamanho do arquivo
		seja maior que o total da memória do sistema

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
	*/

	/* 	Dessa forma o go usa um buffer de um tamanho especificado pelo buffer := make([]byte, QTD_BYTES)
	sendo QTD_BYTES a quantidade total de bytes que será lido or cada iteração que será feita pelo
	reader, assim permitindo ler arquivos grandes */

	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
