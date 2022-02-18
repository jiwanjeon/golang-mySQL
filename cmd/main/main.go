package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jiwanjeon/golang-mySQL/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRouts(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
