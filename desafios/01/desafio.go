package main

import "fmt"

func main() {
	for n := 33; n < 123; n++ {
		fmt.Printf("%d - %c\n", n, n)
	}
}
