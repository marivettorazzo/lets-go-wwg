package main

import (
	"fmt"
)

func main() {
	mostraNomesNaOrdem()
	listaDeMercado()
}

func mostraNomesNaOrdem() {
	meuArray := []string{"Pope", "Biguis", "Pibizinho"}
	for indice, valor := range meuArray {
		fmt.Printf("%d - %s\n", indice, valor)
	}
}
func listaDeMercado() {
	listaDeMercado := []string{"Tomate", "Abacate", "Banana", "Sabonete", "azeite"}
	for indice, valor := range listaDeMercado {
		indice++
		fmt.Printf("%d - %s\n", indice, valor)
	}
}
