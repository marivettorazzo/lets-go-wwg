package main

import "fmt"

func main() {

	mostraSequencia()

	mostraHoras()

}
func mostraHoras() {
	for hora := 0; hora < 3; hora++ {
		for minuto := 0; minuto <= 59; minuto++ {
			for segundos := 0; segundos <= 59; segundos++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundos)
			}

		}
	}
}
func mostraSequencia() {
	for i := 13; i <= 27; i++ {
		fmt.Println(i)
	}
}
