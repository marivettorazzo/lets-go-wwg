package main

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
	cor   string
}

func main() {

	pessoa1 := Pessoa{
		nome:  "Fernanda",
		idade: 30,
		cor:   "Azul",
	}
	pessoa2 := Pessoa{
		nome:  "Marta",
		idade: 20,
		cor:   "Rosa",
	}
	pessoa3 := Pessoa{
		nome:  "Mariana",
		idade: 31,
		cor:   "Roxo",
	}

	fmt.Println("Nome:", pessoa1.nome, ",", "Idade:", pessoa1.idade, ",", "Cor favorita:", pessoa1.cor)
	fmt.Println("Nome:", pessoa2.nome, ",", "Idade:", pessoa2.idade, ",", "Cor favorita:", pessoa2.cor)
	fmt.Println("Nome:", pessoa3.nome, ",", "Idade:", pessoa3.idade, ",", "Cor favorita:", pessoa3.cor)
	fmt.Printf("A %s tem %d anos e sua cor favorita é %s.\n", pessoa3.nome, pessoa3.idade, pessoa3.cor)
	idade := 10

	//if idade >= 18 && idade < 60{
	//fmt.Printf("Você é um jovem adulto com a idade de %d\n",idade)
	//return
	//} else if  idade > 60 {
	//	fmt.Printf("Você é idoso com a idade de %d\n",idade)
	//	return
	//	} else if idade >= 0 && idade < 18{
	//	fmt.Printf("Você é uma criança com a idade de %d\n",idade)
	//	return
	//	}
	//	fmt.Println("Informe uma idade  válida\n")

	switch {
	case idade >= 18 && idade < 60:
		fmt.Printf("Você é um jovem adulto com a idade de %d anos\n", idade)
	case idade > 60:
		fmt.Printf("Você é idoso com a idade de %d anos\n", idade)
	case idade >= 0 && idade < 18:
		fmt.Printf("Você é uma criança com a idade de %d anos\n", idade)
	default:
		fmt.Println("Informe uma idade  válida")

	}
	hora := 23.59
	switch {
	case hora > 12 && hora < 18:
		fmt.Printf("Período de %f horas é tarde", hora)
	case hora < 12 && hora > 5:
		fmt.Printf("Período de %f horas é manhã", hora)
	case hora <= 5:
		fmt.Printf("Período de %f horas é madrugada", hora)
	case hora >= 18 && hora <= 24:
		fmt.Printf("Período de %f horas é noite", hora)
	default:
		fmt.Println("Informe uma hora  válida")

	}

}
