package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sorvetes  []string
}

func main() {
	pessoas := map[string]pessoa{
		"Sales": {
			nome:      "Guilherme",
			sobrenome: "Sales",
			sorvetes:  []string{"chocolate", "baunilha", "ferreiro roche"},
		},
		"Marquezini": {
			nome:      "Gabrielly",
			sobrenome: "Marquezini",
			sorvetes:  []string{"flocos", "coco", "napolitano"},
		},
	}
	for _, v := range pessoas {
		fmt.Println("Os sabores favoritos de", v.nome, v.sobrenome, ":")
		for _, v := range v.sorvetes {
			fmt.Print("\t--", v, "\n")
		}
	}
}
