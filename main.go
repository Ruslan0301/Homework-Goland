package main

import (
	"github.com/gorilla/mux"
	"github.com/upper/db/v4/adapter/postgresql"
	"net/http"
)

var settings = postgresql.ConnectionURL{
	Database: `Events_base`,
	Host:     `localhost:8080`,
	User:     `postgres`,
	Password: `060378so`,
}

type Event struct {
	ID                uint   `db:"id,omitempty"`
	EventsName        string `db:"events_name"`
	EventsDescription string `db:"events_description"`
	City              string `db:"events_city"`
	EventsAddress     string `db:"events_address"`
}

func main() {
	handleFunc()
	//
}

func handleFunc() {
	//Роутер
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", mainPage).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_in_base", save_in_base).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", see_more).Methods("GET")
	rtr.HandleFunc("/search", search).Methods("GET", "POST")

	//http.HandleFunc
	http.Handle("/", rtr)
	http.ListenAndServe("localhost:8089", nil)

}
