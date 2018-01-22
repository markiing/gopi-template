package book

import (
	"net/http"
	"gopi-template/connection"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Book struct{
	Nome	string
	Isbn	string
}

func ShowAll(w http.ResponseWriter, r *http.Request){
	var book []Book;
	connection.Find("books",bson.M{"nome":"Livro Teste"}).All(&book);
	json.NewEncoder(w).Encode(book)
}