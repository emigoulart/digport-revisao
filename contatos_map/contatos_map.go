package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func main() {
	contacts := make(map[string]Contact)

	// Adicionando contatos ao mapa
	contacts["John"] = Contact{Name: "John Doe", Phone: "123456789"}
	contacts["Jane"] = Contact{Name: "Jane Smith", Phone: "987654321"}

	// Imprimindo os contatos existentes
	fmt.Println("Contatos existentes:")
	for _, contact := range contacts {
		fmt.Println("Nome:", contact.Name)
		fmt.Println("Telefone:", contact.Phone)
		fmt.Println()
	}

	// Adicionando novos contatos
	contacts["Bob"] = Contact{Name: "Bob Johnson", Phone: "555555555"}
	contacts["Alice"] = Contact{Name: "Alice Brown", Phone: "999999999"}

	var get_contato string
	fmt.Printf("digite o nome do contato que deseja encontrar: ")
	fmt.Scan(&get_contato)

	v, encontrada := contacts[get_contato] // valor e a variavel boolean para verificar o contato pela chave string
	if encontrada {
		fmt.Printf("O detalhes do contato %s informado é: %s\n", get_contato, v)
	} else {
		fmt.Printf("Contato %s não encontrado\n", get_contato)

	}

	// Imprimindo os contatos atualizados
	fmt.Println("Contatos atualizados:")
	for _, contact := range contacts {
		fmt.Println("Nome:", contact.Name)
		fmt.Println("Telefone:", contact.Phone)
		fmt.Println()
	}
}

//Neste exemplo, criamos um mapa chamado `contacts`, onde a chave é uma string representando
//o nome do contato e o valor é uma estrutura `Contact` que contém os detalhes do contato (nome e telefone).
//Em seguida, adicionamos alguns contatos ao mapa e os imprimimos. Depois, adicionamos mais contatos
//e imprimimos novamente para mostrar os contatos atualizados.
