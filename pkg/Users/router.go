package Users

import (
	"github.com/go-chi/chi"
)

func UserRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/getUser", getHandler)
	router.Post("/postUser", postHandler)

	return router
}
