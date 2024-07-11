package contato

import "fmt"

type Contato struct {
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

var contatos []Contato = []Contato{}

func ListaContatos() []Contato {
	contatosList := []Contato{
		{Nome: "Maria", Telefone: "987654321", Email: "maria@example.com"},
		{Nome: "João", Telefone: "123456789", Email: "joao@example.com"},
	}
	if len(contatos) > 0 {
		contatosList = contatos // retorna a lista atualizada de contatos, caso adicionado um novo contato
	}
	return contatosList

}

func AddContato(contato Contato) bool {
	if contatoExiste(ListaContatos(), contato.Nome) {
		fmt.Printf("Contato de nome %s já existe", contato.Nome)
		return false
	}

	contatos = append(ListaContatos(), contato) //adiciona o contato enviado via post a lista de contatos
	return true
}

func contatoExiste(contatos []Contato, nome string) bool {
	for _, contato := range contatos {
		if contato.Nome == nome {
			return true
		}
	}
	return false
}
