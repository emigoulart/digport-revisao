package main

import (
	"fmt"

	"gopkg.in/Nhanderu/brdoc.v1"
)

func main() {
	var cpf string
	fmt.Print("Enter CPF: ")
	fmt.Scan(&cpf)

	if brdoc.IsCPF(cpf) {
		fmt.Println("Valid CPF")
	} else {
		fmt.Println("Invalid CPF")
	}
}
