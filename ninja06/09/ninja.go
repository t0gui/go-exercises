package main

import "fmt"

func nomeCompleto(nome string, sobrenome string, adjetivo func() string) string {
	return adjetivo() + " " + nome + " " + sobrenome
}

func bonito() string {
	return "voce é lindo"
}

func main() {
	elogiar := nomeCompleto("Guilherme", "Sales", bonito)
	fmt.Println(elogiar)
}
