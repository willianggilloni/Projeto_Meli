package main

import (
	"fmt"
	"log"
	"net/http"
	"testMeli/src/config"
	"testMeli/src/router"
)

func main() {

	println("Rodando api")
	config.Carregar()
	r := router.Gerar()
	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
