package main

import (
	"encoding/json"
	"net/http"

	"github.com/emigoulart/digport-revisao/contato"
	"github.com/gorilla/mux"
)

func Rotas() *mux.Router {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/lista/contatos", HandleContato).Methods("GET")
	rotas.HandleFunc("/add/contato", HandleAddContato).Methods("POST")
	return rotas
}

func HandleContato(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contato.ListaContatos())
}

func HandleAddContato(w http.ResponseWriter, rq *http.Request) {
	var conts []contato.Contato                    //slice de contatos
	err := json.NewDecoder(rq.Body).Decode(&conts) //decode json into conts
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
		contato.AddContato(cont)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

}

// referencia : https://aprendagolang.com.br/?s=mux
// para baixar a lib go get -u github.com/gorilla/mux
// http://localhost:8085/lista/contatos
