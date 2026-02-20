package main

import "fmt"

type quadrado struct {
	lado float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 2 * 3.14159 * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{lado: 5}
	c := circulo{raio: 3}
	// Demonstrando a área do Quadrado
	fmt.Printf("Quadrado com lado %.2f\n", q.lado)
	fmt.Printf("Área do Quadrado: %.2f\n\n", info(q))

	// Demonstrando a área do Círculo
	fmt.Printf("Círculo com raio %.2f\n", c.raio)
	fmt.Printf("Área do Círculo: %.2f\n", info(c))
}
