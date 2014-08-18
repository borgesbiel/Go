package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{}
	fmt.Println("Stack created with size: ", pilha.Tamanho())
	fmt.Println("Empty? ", pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("End")

	fmt.Println("Size after stack 4 values: ", pilha.Tamanho())
	fmt.Println("Empty? ", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("Unstacking ", v)
		fmt.Println("Size: ", pilha.Tamanho())
		fmt.Println("Empty? ", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Empty stack!")
	}

	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]

	return valor, nil
}
