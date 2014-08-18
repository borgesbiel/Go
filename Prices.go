package main

import "fmt"

func PrecoFinal(precoCusto float64) (precoReal, precoDolar float64) {
	fatorLucro := 1.33
	taxaCovercao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaCovercao

	return
}

func main() {
	precoReal, precoDolar := PrecoFinal(34.99)
	fmt.Printf("Dolar: %.2f\n"+"Real: %.2f\n", precoDolar, precoReal)
}
