package main

import "fmt"

const (
	x        = 10
	y string = "Eu sou"
)

func main() {
	fmt.Printf("%v -> %T\n%v -> %T\n", x, x, y, y)
}
