package main //PRINCIPAL PACOTE DA APLICAÇÃO

import (
	"gopi-template/routes"
	"log"
	"net/http"
)


func main() {
	router := routes.NewRouter();
	log.Fatal(http.ListenAndServe(":7076",router))
}
