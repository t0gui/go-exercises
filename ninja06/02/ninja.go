package main

import "fmt"

func main() {
	numeros1 := []int{1, 2, 3, 4, 5}
	fmt.Println(somar(numeros1...))

	numeros2 := []int{6, 7, 8, 9, 10}
	fmt.Println(somaSlice(numeros2))
}

func somar(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

func somaSlice(numeros []int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}
