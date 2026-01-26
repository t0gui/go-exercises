package main

import "fmt"

func main() {
	nasci := 1994
	atual := 2026
	for {
		if nasci > atual {
			break
		}
		fmt.Println(nasci)
		nasci++
	}
}
