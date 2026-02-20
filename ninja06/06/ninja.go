package main

import "fmt"

func main() {
	func(nome string) {
		fmt.Printf("%s é um programador\n", nome)
	}("Guilherme Sales")
}
