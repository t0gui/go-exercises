package main

import "fmt"

func main() {
	familia := map[string][]string{
		"guilherme": {
			"hackear", "programar", "estudar",
		},
		"gaby": {
			"estudar", "ler", "projetar",
		},
		"marineuza": {
			"dormir", "costurar", "fofocar",
		},
	}
	for nome, hobbie := range familia {
		fmt.Printf("%v\n", nome)
		for i, hoby := range hobbie {
			fmt.Printf("\t%v- %v\n", i, hoby)
		}
	}
}
