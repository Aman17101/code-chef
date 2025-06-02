package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Salary int
}

type errr struct {
	Status int
	msg    string
}

var (
	users = map[(string)]User{}
)

func main() {

	http.HandleFunc("/createUser", addUser)
	fmt.Println("server started")

	log.Fatalf("server start nhi hua %v ", http.ListenAndServe(":8000", nil))
}
func addUser(p http.ResponseWriter, q *http.Request) {
	if q.Method != "POST" {
		p.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := User{}
	err := json.NewDecoder(q.Body).Decode(&user)
	if err != nil {
		p.WriteHeader(http.StatusBadRequest)
		errrr := errr{
			Status: http.StatusBadRequest,
			msg:    "Invalid JSON"+err.Error() ,
		}
		json.NewEncoder(p).Encode(errrr)

		return
	}
	users[user.Name] = user

	p.WriteHeader(http.StatusCreated)

	json.NewEncoder(p).Encode(user)

	return

}
