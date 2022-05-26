package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/upper/db/v4/adapter/postgresql"
	"html/template"
	"log"
	"net/http"
)

var eventsList = []Event{}

var showEventDescription = Event{}

func see_more(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tmpl, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Connect to base.
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	fmt.Println("Ok")
	defer db.Close()

	//Вибірка даних
	res, err := db.SQL().Query("SELECT * FROM event WHERE events_id = ?", vars["id"])
	if err != nil {
		log.Fatal("Select: ", err)
	}

	//Значення пустої структури
	showEventDescription = Event{}

	//Цикл
	for res.Next() {
		var post Event
		err = res.Scan(&post.ID, &post.EventsName, &post.EventsDescription, &post.City, &post.EventsAddress)
		if err != nil {
			log.Fatal("Row: ", err)
		}
		showEventDescription = post
	}
	tmpl.ExecuteTemplate(w, "show", showEventDescription)
}

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create.html",
		"templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "create", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "home", eventsList)
}

func search(w http.ResponseWriter, r *http.Request) {
	//Отримаємо значення поля пошуку (місто)
	city := Event{
		City: r.FormValue("select_city"),
	}

	tmpl, err := template.ParseFiles("templates/home.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//base connect
	db, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	fmt.Println("Ok")
	defer db.Close()

	//Вибірка даних
	res, err := db.SQL().Query("SELECT * FROM event WHERE events_city = ?", city.City)
	if err != nil {
		log.Fatal("Select: ", err)
	}

	//Чистим список івентів
	eventsList = []Event{}

	//Цикл
	for res.Next() {
		var post Event
		err = res.Scan(&post.ID, &post.EventsName, &post.EventsDescription, &post.City, &post.EventsAddress)
		if err != nil {
			log.Fatal("Row: ", err)
		}
		//fmt.Println(fmt.Sprintf("Event: %s with id: %d", post.EventsName, post.ID))
		eventsList = append(eventsList, post)
	}
	tmpl.ExecuteTemplate(w, "home", eventsList)
}
