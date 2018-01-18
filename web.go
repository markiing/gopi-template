package main //PRINCIPAL PACOTE DA APLICAÇÃO

import (
	"fmt"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4567", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var p = Page{Title: "Meu Titulo", Content: "Conteudo teste"}
	fmt.Println(p)
	fmt.Fprintf(w, "Hi there, I love %s!", p)
}
