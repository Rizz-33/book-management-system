package main

import (
	"log"
	"net/http"

	"github.com/Rizz-33/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Learning CI/CD Pipeline

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
