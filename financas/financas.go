package main

import "fmt"

func main() {
	despesas := map[string]int{"mercado": 1500, "aluguel": 2200, "eletricidade": 3500}

	total := 0

	var get_despesa string
	fmt.Printf("digite a despesa: ")
	fmt.Scan(&get_despesa)

	v, encontrada := despesas[get_despesa] // valor e a variavel boolean para verificar a despesa pela chave string
	if encontrada {
		fmt.Printf("O valor da despesa %s informada é: %d\n", get_despesa, v)
	} else {
		fmt.Printf("Despesa %s não encontrada\n", get_despesa)

	}
	for despesa, valor := range despesas {
		total = total + valor
		fmt.Printf("A despesa de %s é %d\n", despesa, valor)
	}

	if total > 2000 {
		fmt.Printf("O total das despesas %d ultrapassou o seu orçamento \n", total)
	}
	fmt.Printf("O total das despesas é %d\n", total)
}
