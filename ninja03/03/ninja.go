package main

import "fmt"

func main() {
	nasci := 1994
	atual := 2027
	for nasci <= atual {
		fmt.Printf("ano - %v\n", nasci)
		nasci++
	}
}
