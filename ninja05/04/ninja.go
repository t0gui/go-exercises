package main

import (
	"fmt"
)

func main() {
	escola := struct {
		nome   string
		alunos []string
		notas  map[string]float64
	}{
		nome:   "hackinagem",
		alunos: []string{"Guilherme", "Gabe", "Marineuza"},
		notas: map[string]float64{
			"Guilherme": 8.5,
			"Gabe":      7.8,
			"Marineuza": 9.2,
		},
	}
	fmt.Printf("Escola: %s\n", escola.nome)
	fmt.Printf("Total de alunos: %d\n\n", len(escola.alunos))
}
