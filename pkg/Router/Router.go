package Router

import (
	"Golang_cql/pkg/Users"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/api/users", Users.UserRoutes())
	fmt.Println("server is listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
