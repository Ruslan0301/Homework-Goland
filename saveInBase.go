package main

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"net/http"
)

func save_in_base(w http.ResponseWriter, r *http.Request) {

	event := Event{
		EventsName:        r.FormValue("event_name"),
		City:              r.FormValue("city"),
		EventsDescription: r.FormValue("description"),
		EventsAddress:     r.FormValue("address"),
	}

	// Connect to base.
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	fmt.Println("Ok")
	defer db.Close()

	if event.EventsName != "" || event.City != "" || event.EventsAddress != "" {
		//add data
		_, err = db.SQL().InsertInto("event").Values(event).Exec()
		if err != nil {
			log.Fatal("Query: ", err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
	}

}
