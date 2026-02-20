package main

import "fmt"

func saudacao(nome string) func(string) string {
	return func(texto string) string {
		return "Bem vindo " + texto + " " + nome
	}
}

func main() {
	formal := saudacao("Guilherme")
	fmt.Println(formal("Sr."))
}
