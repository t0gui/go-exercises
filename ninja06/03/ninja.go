package main

import "fmt"

func main() {
	defer fmt.Println("Os primeiros serao os ultimos")
	fmt.Println("Os ultimos serao os primeiros")
}
