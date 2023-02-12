package main

import (
	"fmt"
	"go-proxy/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/", routes.ProxyHandler)

	http.HandleFunc("/finalizaRequisicao/", routes.FinalizaHandler)

	fmt.Println("iniciando proxy na porta 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
