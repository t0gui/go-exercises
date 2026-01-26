package main

import "fmt"

func main() {
	pizzas := []string{"peperoni", "marguerita", "calabresa", "moda"}
	for n := 0; n < len(pizzas); n++ {
		fmt.Println(pizzas[n])
	}
	fmt.Println(pizzas[:])
}
