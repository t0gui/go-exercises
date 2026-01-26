package main

import "fmt"

func main() {
	a := 10 == 10
	b := 10 != 5
	c := 30 <= 20
	d := 10 < 5
	e := 10 >= 5
	f := 10 > 5

	fmt.Println(a, b, c, d, e, f)
}
