package main

import (
	"desafio/routes"
	"net/http"
)

//MAIN
func main() {
	routes.CarregaRotasAPI()
	http.ListenAndServe(":80", nil)
}
