package main

import (
	"fmt"
	infra_database "golang/projects/go-hexagonal/infra/database"
	infra_secrets "golang/projects/go-hexagonal/infra/secrets"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	infra_secrets.LoadEnv()
	infra_database.LoadDatabase()
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", getPost).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	port := os.Getenv("PORT")

	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)

	if err != nil {
		log.Fatal(err)
	}
}
