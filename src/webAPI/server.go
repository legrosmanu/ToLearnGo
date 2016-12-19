package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ErrorToLearnGo : error for this test
type ErrorToLearnGo struct {
	httpCode int
	message  string
	exists   bool
}

// User : the oject that I want to send on JSON by his web API
type User struct {
	firstname, lastname string
}

func getUserByID(id string) (user User, err ErrorToLearnGo) {
	if id == "legrosmanu" {
		user.firstname = "Emmanuel"
		user.lastname = "LEGROS"
	} else {
		err = ErrorToLearnGo{404, "User nof found", true}
	}
	return
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users/{id}", getUser)

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err.exists {
		if err.httpCode == 404 {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintln(w)
	} else {
		fmt.Fprintln(w, user)
	}
}
