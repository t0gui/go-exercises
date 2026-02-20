package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) apresentar() string {
	return fmt.Sprintf("Ola meu nome é %s %s e eu tenho %d anos", p.nome, p.sobrenome, p.idade)
}

func main() {
	eu := pessoa{nome: "Guilherme", sobrenome: "Sales", idade: 31}
	fmt.Println(eu.apresentar())

	amor := pessoa{nome: "Gabe", sobrenome: "Marquezini", idade: 27}
	fmt.Println(amor.apresentar())
}
