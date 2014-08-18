package main

import (
	"fmt"
)

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]

	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista) - 1)
}

func main() {
	lista := ListaGenerica{1, "Coffee", 42, true, 23, "Ball", 3.14, false}

	fmt.Printf("List:\n%v\n\n", lista)

	fmt.Printf(
		"Removing at start: %v, after remove:\n%v\n",
		lista.RemoverInicio(), lista)
	fmt.Printf(
		"Removing at end: %v, after remove:\n%v\n",
		lista.RemoverFim(), lista)
	fmt.Printf(
		"Removing from index 3: %v, after remove:\n%v\n",
		lista.RemoverIndice(3), lista)
	fmt.Printf(
		"Removing from index 0: %v, after remove:\n%v\n",
		lista.RemoverIndice(0), lista)
	fmt.Printf(
		"Removing from last index: %v, after remove:\n%v\n",
		lista.RemoverIndice(len(lista)-1), lista)
}
