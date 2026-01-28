package main

import "fmt"

func main() {
	pessoas := [][]string{
		{"Guilherme", "Sales", "Hackear"},
		{"Gabe", "Marquezini", "Arquitetar"},
		{"Marineuza", "Pereira", "Costurar"},
	}
	for _, pessoa := range pessoas {
		fmt.Printf("%s %s - %s\n", pessoa[0], pessoa[1], pessoa[2])
	}
}
