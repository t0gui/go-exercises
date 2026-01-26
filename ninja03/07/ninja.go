package main

import "fmt"

func main() {
	niveldeDor := 4
	switch {
	case niveldeDor == 1:
		fmt.Println("Nao ta doendo nada")
	case niveldeDor == 2:
		fmt.Println("Nem pense em chorar")
	case niveldeDor == 3:
		fmt.Println("Estou ficando preocupado")
	default:
		fmt.Println("Bora pro medico")

	}
}
