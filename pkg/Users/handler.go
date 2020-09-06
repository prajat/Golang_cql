package Users

import (
	"encoding/json"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("you got user 1")
}
func postHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("you update the user 1")
}
