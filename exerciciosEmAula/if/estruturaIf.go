package main

import "fmt"

func main() {
	fmt.Println((podeVotar(1990)))
}

func podeVotar(anoNascimento int) string {

	anoAtual := 2021

	if anoAtual-anoNascimento > 16 {

		return "Você já pode votar."

	}
	return "Você ainda não pode votar."

}
