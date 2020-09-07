package Users

import (
	"Golang_cql/Database"
	"encoding/json"
	"fmt"
	"net/http"
)

var user User

func getHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("you got user 1")
}
func postHandler(w http.ResponseWriter, r *http.Request) {

	//json.NewEncoder(w).Encode("you update the user 1")
	json.NewDecoder(r.Body).Decode(&user)
	addUser(&user)
	fmt.Println("just addded user to the database")
	json.NewEncoder(w).Encode("you just added the user 1 to cassandra DB")
}
func addUser(user *User) {

	query := `INSERT INTO users (first_name,last_name,email) VALUES (?,?,?)`
	Database.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}
