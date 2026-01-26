package main

import "fmt"

func main() {
	switch esporteFavorito := "skate"; esporteFavorito {
	case "futebol":
		fmt.Printf("Meu esporte favorito é %v\n", esporteFavorito)
	case "sinuca":
		fmt.Printf("Meu esporte favorito é %v\n", esporteFavorito)
	case "skate":
		fmt.Printf("Meu esporte favorito é %v\n", esporteFavorito)

	}
}
