package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range slice {
		fmt.Println(i, n)
	}
	fmt.Printf("Tipo --> %T\n", slice)
}
