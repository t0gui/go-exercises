package main

import "fmt"

func contar(inicial int) func() int {
	numeroAtual := inicial
	return func() int {
		numeroAtual++
		return numeroAtual
	}
}

func main() {
	numero := contar(5)
	fmt.Println(numero())
}
