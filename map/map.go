package main

import "fmt"

func main() {
	funcionarios := map[string]int{"Pedro": 1500, "João": 2200, "Maria": 3500}
	delete(funcionarios, "Pedro")
	funcionarios["Peter"] = 5000

	// func := make(map[string]int)
	// func1 := map[string]int{}
	// func1["Emilene"] = 4000

	for nome, funcionario := range funcionarios {
		fmt.Printf("O salario de %s é %d\n", nome, funcionario)
	}

	// omitindo a chave nome
	for _, salario := range funcionarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
