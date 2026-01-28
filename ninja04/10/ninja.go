package main

import "fmt"

func main() {
	notas := map[string][4]float64{
		"Guilherme": {8.5, 9.0, 7.5, 8.0},  // 4 notas
		"Gabe":      {6.0, 7.0, 8.0, 7.5},  // 4 notas
		"Marineuza": {9.5, 9.0, 10.0, 9.8}, // 4 notas
	}
	delete(notas, "Marineuza")
	for name, nota := range notas {
		fmt.Printf("%s\n", name)
		for i, n := range nota {
			fmt.Printf("ID %d -> Notas %.2f\n", i, n)
		}
	}
}
