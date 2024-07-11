package main

import (
	"net/http"
)

func main() {
	r := Rotas()
	http.ListenAndServe(":8085", r)
}
