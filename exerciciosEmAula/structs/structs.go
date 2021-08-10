package main

import "fmt"

type Alguem struct {
	nome  string
	idade int
	cor   string
}

func main() {
	var nome string
	var idade int
	var cor string
	fmt.Print("Digite, seu nome, sua idade e sua cor favorita :")
	fmt.Scanf("%s", &nome)
	fmt.Scanf("%d", &idade)
	fmt.Scanf("%s", &cor)
	pessoa := Alguem{nome, idade, cor}
	fmt.Println(pessoa)

}
