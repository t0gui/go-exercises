package main

import "fmt"

func main() {
	for n := 65; n <= 90; n++ {
		fmt.Println(n)
		for i := 0; i < 3; i++ {
			fmt.Printf("%U\n", n)
		}
	}
}
