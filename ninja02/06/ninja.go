package main

import "fmt"

func main() {
	const (
		_ = iota + 1993
		a
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
