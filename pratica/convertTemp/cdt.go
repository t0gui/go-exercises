package main

import "fmt"

func main() {
	fmt.Println("=== Conversor de Temperatura ===")
	fmt.Println("1 - Celsius para Fahrenheit e Kelvin")
	fmt.Println("2 - Fahrenheit para Celsius e Kelvin")
	fmt.Println("3- Kelvin para Fahrenheit e Celsius")
	fmt.Print("Digite a opcao desejada: ")

	var opcao int

	_, err1 := fmt.Scanln(&opcao)
	if err1 != nil {
		fmt.Println("Erro ao ler a opcao!")
	}

	fmt.Print("Digite a temperatura: ")
	var valor float64

	_, err2 := fmt.Scanln(&valor)
	if err2 != nil {
		fmt.Println("Erro de valor")
	}

	switch opcao {
	case 1:
		f, k := celsius(valor)
		fmt.Printf("%.2f°C = %.2f°F = %.2f°K\n", valor, f, k)
	case 2:
		c, k := fahrenheit(valor)
		fmt.Printf("%.2f°F = %.2f°C = %.2f°K \n", valor, c, k)
	case 3:
		f, c := kelvin(valor)
		fmt.Printf("%.2f°K = %.2f°C = %.2f°F\n", valor, c, f)
	default:
		fmt.Println("Opcao invalida!")
	}
}

func celsius(celsius float64) (float64, float64) {
	f := (celsius * 9 / 5) + 32
	k := (celsius + 273.15)
	return f, k
}

func fahrenheit(fahrenheit float64) (float64, float64) {
	c := (fahrenheit - 32) * 5 / 9
	k := (fahrenheit-32)*5/9 + 273.15
	return c, k
}

func kelvin(kelvin float64) (float64, float64) {
	c := kelvin - 273.15
	f := (kelvin-273.15)*9/5 + 32
	return c, f
}
