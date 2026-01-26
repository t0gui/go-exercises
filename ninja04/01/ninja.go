package main

import "fmt"

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	for i, v := range num {
		fmt.Printf("Indice %d: %d\n", i, v)
	}
	fmt.Printf("Tipo: %T\n", num)
}
