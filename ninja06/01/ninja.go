package main

import "fmt"

func main() {
	fmt.Println(umafuncao())
	fmt.Println(outrafuncao())
}

func umafuncao() int {
	return 10
}

func outrafuncao() (string, int) {
	return "Eu sou", 10
}
