package main

import "fmt"

type veiculo struct {
	porta int
	cor   string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	s10 := caminhonete{veiculo: veiculo{porta: 4, cor: "branco"}, tracaoNasQuatro: true}
	civic := sedan{veiculo{4, "prata"}, false}
	fmt.Println("\tCaracteristicas da S10")
	fmt.Printf("Cor: %v\nPortas: %v\nTem tracao nas 4: %v\n", s10.cor, s10.porta, s10.tracaoNasQuatro)
	fmt.Println("\tCaracteristicas do Civic")
	fmt.Printf("Cor:%v\nPortas:%v\nLuxo:%v\n", civic.cor, civic.porta, civic.modeloLuxo)
}
