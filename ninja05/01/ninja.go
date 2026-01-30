package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   []string
}

func main() {
	p1 := pessoa{
		Nome:      "Guilherme",
		Sobrenome: "Sales",
		sabores:   []string{"Napolitano", "Baunilha", "Chocolate"},
	}
	p2 := pessoa{
		Nome:      "Gabielly",
		Sobrenome: "Marquezini",
		sabores:   []string{"Flocos", "Maracuja", "Ouro Branco"},
	}
	fmt.Println(p1.Nome, p1.Sobrenome)
	for _, p := range p1.sabores {
		fmt.Println("\t", p)
	}
	fmt.Println(p2.Nome, p2.Sobrenome)
	for _, p := range p2.sabores {
		fmt.Println("\t", p)
	}
}
