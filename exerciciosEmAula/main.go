package main

import (
	"fmt"
	"strconv"
	"time"
)

// exercicios extras
func main() {
	qualIdade()
	fmt.Println("-----------------")
	quilometragem()
	fmt.Println("-----------------")
	times()
	fmt.Println("-----------------")
	addJogadores()
	fmt.Println("-----------------")
	qteVezesQueApareceNoMapa()
	fmt.Println("-----------------")
	maior(25, 50, 15)
	fmt.Println("-----------------")
	multiploDe()
	fmt.Println("-----------------")
	listaDeMercadoComForTrad()
	fmt.Println("-----------------")
	mostraSequencia(10)
	fmt.Println("-----------------")
	validacao()

}

func qualIdade() {
	fmt.Print("Enter your birth year: ")
	var anoNascimento int
	fmt.Scanln(&anoNascimento)
	anoAtual := time.Now()
	idade := anoAtual.Year() - anoNascimento
	fmt.Printf("Sua iade é de : %d\n", idade)

}

func quilometragem() {
	quilometros := 150
	fmt.Println(quilometros)
}

func times() {
	amarelo := [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := [4]string{"Helena", "Jonas", "José ", "Juliana"}
	fmt.Println("time: Amarelo", amarelo, "time vermelho: ", vermelho)

}
func addJogadores() {
	amarelo := [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := []string{"Helena", "Jonas", "José ", "Juliana"}
	vermelho = append(vermelho, "luiz Inácio")
	fmt.Println("time: Amarelo", amarelo, "time vermelho: ", vermelho)

}
func qteVezesQueApareceNoMapa() {
	paises := map[string]string{"saoPaulo": "Brasil", "rioDeJaneiro": "Brasil", "paris": "Franca", "cannes": "Franca", "toronto": "Canada", "vancouver": "Canada"}

	var sPaises []string
	var sRepetidos []string
	var array []int

	for _, pais := range paises {
		sPaises = append(sPaises, pais)
	}

	for _, p := range sPaises {
		qtd := find(sPaises, p)
		jaInserido := find(sRepetidos, p)

		if qtd > 1 && jaInserido == 0 {
			sRepetidos = append(sRepetidos, p)
			array = append(array, qtd)

		}
	}
	fmt.Println(array)
}

func find(item1 []string, valor string) int {
	total := 1
	for _, item := range item1 {
		if item == valor {
			total++
			return total
		}
	}
	return 0
}
func maior(a int, b int, c int) {
	if a > b && a > c {
		fmt.Printf("%d", a)
	} else if b > a && b > c {
		fmt.Printf("%d \n", b)
	} else {
		fmt.Printf("%d \n", c)
	}
}
func multiploDe() {
	fmt.Print("Enter with number: ")
	var num int
	fmt.Scanln(&num)

	option := num%2 == 0 && num != 0
	option2 := num == 0
	option3 := num%3 == 0

	switch {
	case option:
		fmt.Printf("o número %d é divisivel por 2 \n", num)
	case option2:
		fmt.Printf("o número é %d \n", num)
	case option3:
		fmt.Printf("O número %d é divisivel por 3 \n", num)
	default:
		fmt.Printf("O numero %d não se aplica a nenhuma das condições \n", num)

	}
}
func listaDeMercadoComForTrad() {

	listaDeMercado := []string{"Tomate", "Abacate", "Banana", "Sabonete", "azeite"}
	for i := 0; i < len(listaDeMercado); i++ {
		item := listaDeMercado[i]
		fmt.Printf("%s \n ", item)
	}

}
func mostraSequencia(num int) {
	temp := " "

	for i := 0; i < num; i++ {

		temp += strconv.Itoa(i) + " "
		fmt.Println(temp)

	}

}

func validacao() {
	var num int
	fmt.Print("Enter number:")
	fmt.Scan(&num)
	for valida(num) != 0 {
		var par int
		fmt.Print("Enter number:")
		fmt.Scan(&par)
		valida(par)
		if valida(par) == 0 {
			fmt.Print("Agora sim podemos dividir igualmente entre mim e você!")
			return
		}
	}
	fmt.Print("Agora sim podemos dividir igualmente entre mim e você!")

}
func valida(num int) int {
	a := num % 2
	return a
}
