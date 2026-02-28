package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	p.idade++
}

func main() {
	p1 := pessoa{
		nome:      "Guilherme",
		sobrenome: "Sales",
		idade:     31,
	}
	fmt.Println(p1)
	mudeMe(&p1)
	fmt.Println(p1)
}
