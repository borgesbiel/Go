package main

import (
	"fmt"
	"os"
)

func CriarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf("%s/%s.$s", dirBase, nome, "txt")

		arq, err := os.Create(caminhoArquivo)

		defer arq.Close()

		if err != nil {
			fmt.Println("Error creating file %s: %v\n", nome, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created.", arq.Name())
	}
}

func main() {
	tmp := os.TempDir()
	CriarArquivos(tmp)
	CriarArquivos(tmp, "teste1")
	CriarArquivos(tmp, "teste2", "teste3", "teste4")
}
