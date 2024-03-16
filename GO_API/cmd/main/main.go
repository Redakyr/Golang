package main

/*
------------ Required Packages ------------

github.com/jinzhu/gorm
github.com/jinzhu/gorm/dialects/mysql
github.com/gorilla/mux

--------------------------------------------
*/

import (
	"github.com/Redakyr/Golang/GO_API/pkg/routes/"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.registerAgentRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
