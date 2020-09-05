package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	fmt.Println("starting server....")

	router := chi.NewRouter()
	router.Get("/api/getExample", getHandler)
	router.Post("/api/postExample", postHandler)

	fmt.Println("server is listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
func getHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("you got me")
}
func postHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("you sent me the post request")
}
