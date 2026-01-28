package main

import "fmt"

func main() {
	br := make([]string, 0, 26)
	br = []string{
		"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
		"Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins",
	}
	fmt.Println("Tamanho atual -> ", len(br))
	fmt.Println("Capacidade total ->", cap(br))
	for i := 0; i < len(br); i++ {
		fmt.Println(br[i])
	}
}
