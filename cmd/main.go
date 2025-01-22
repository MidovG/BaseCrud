package main

import (
	"baseCrud/internal"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", internal.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", internal.GetUserById).Methods("GET")
	r.HandleFunc("/add_user", internal.CreateUser)
	r.HandleFunc("/delete/{id}", internal.DeleteUserById)
	r.HandleFunc("/update/{id}", internal.EditPage).Methods("GET")
	r.HandleFunc("/update/{id}", internal.UpdateUser).Methods("POST")

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
