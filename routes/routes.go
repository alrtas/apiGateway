package routes

import (
	"desafio/controllers"
	"net/http"
)

//CarregaRotasAPI func
func CarregaRotasAPI() {

	http.HandleFunc("/cadastro/geral", controllers.Cadastro_geral)
	http.HandleFunc("/cadastro/financeiro", controllers.Cadastro_financeiro)
	http.HandleFunc("/cadastro/dividas", controllers.Cadastro_dividas)
	http.HandleFunc("/cadastro/transacoes", controllers.Cadastro_transacoes)
	http.HandleFunc("/baseA/dadosGerais/", controllers.BaseA_dadosGerais)
	http.HandleFunc("/baseB/dadosScore", controllers.BaseB_dadosScore)
	http.HandleFunc("/baseC/transacoes", controllers.BaseC_transacoes)

}
