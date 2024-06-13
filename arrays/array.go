package main

import "fmt"

type Usuario struct {
	Nome     string
	Id       string
	Endereco string
	Renda    float64
}

func main() {

	var revistas [4]string
	//var revistas [4]string = [4]string{"Capricho"}
	precos := [5]float64{3.90, 5.60, 10.99, 19.99, 25.00}
	revistas[2] = "Nova"
	fmt.Println(precos)
	fmt.Println(revistas)
	fmt.Println(precos[0])
	// selecionando parte de uma array com slice
	precosSelecionados := precos[0:3]
	fmt.Println(precosSelecionados)
	precosSelecionados = append(precosSelecionados, 35.00)
	fmt.Println(len(revistas))
	fmt.Println(precosSelecionados)

	for i, valor := range precos {
		fmt.Printf("O indice é %d e o valor é %.2f\n", i, valor)
	}
}
