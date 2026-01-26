package main

import "fmt"

func main() {
	nome := "t0gui"
	if nome == "guilherme" {
		fmt.Println("Ola" + nome)
	} else if nome == "t0gui" {
		fmt.Println("Ola " + nome)
	}
}
