package main

import "fmt"

func main() {
	var nome = "Pope"
	var hora = 10

	cumprimento := gereCumprimento(nome, hora)
	fmt.Println(cumprimento)

	soma := somar(10, 7)
	fmt.Println(soma)
}

func gereCumprimento(nome string, hora int) string {
	cumprimento := ""
	switch {
	case hora >= 6 && hora < 12:
		cumprimento = "Bom dia."
	case hora >= 12 && hora < 18:
		cumprimento = "Boa tarde."
	case hora >= 18 && hora < 24:
		cumprimento = "Boa noite."
	case hora >= 0 && hora < 6:
		cumprimento = "Boa madrugada."
	default:
		cumprimento = "Olá!"
	}

	frase := fmt.Sprintf("%s, %s!", cumprimento, nome)
	return frase
}
func somar(a int, b int) int {
	result := a + b
	return result
}
