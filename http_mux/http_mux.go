package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var Contatos []Contato = []Contato{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lista/contatos", HandleContato).Methods("GET")
	r.HandleFunc("/add/contato", HandleAddContato).Methods("POST")
	http.ListenAndServe(":8085", r)
}

func HandleContato(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listaContatos())
}

type Contato struct {
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

func listaContatos() []Contato {
	contatos := []Contato{
		{Nome: "Maria", Telefone: "987654321", Email: "maria@example.com"},
		{Nome: "João", Telefone: "123456789", Email: "joao@example.com"},
	}
	if len(Contatos) > 0 {
		contatos = Contatos
	}

	return contatos

}

func HandleAddContato(w http.ResponseWriter, rq *http.Request) {
	var conts []Contato
	err := json.NewDecoder(rq.Body).Decode(&conts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = json.Marshal(conts)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// add o contato para o slice
	for _, cont := range conts {
		addContato(cont)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	//w.Write(jsonData)
}

func addContato(contato Contato) bool {
	if contatoExiste(listaContatos(), contato.Nome) {
		fmt.Printf("Contato de nome %s já existe", contato.Nome)
		return false
	}

	Contatos = append(listaContatos(), contato)
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

// referencia : https://aprendagolang.com.br/?s=mux
// para baixar a lib go get -u github.com/gorilla/mux
// http://localhost:8085/lista/contatos
