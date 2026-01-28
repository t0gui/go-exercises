package main

import "fmt"

func main() {
	pessoa := make(map[string]string)
	pessoa["Guilherme Sales"] = "Programador"
	pessoa["Gaby Marquezini"] = "Arquiteta"
	pessoa["Marienuza"] = "Costureira"
	i := 0
	for chave, valor := range pessoa {
		fmt.Println(chave, valor, i)
		i++
	}
}
