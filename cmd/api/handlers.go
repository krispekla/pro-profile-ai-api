package main

import (
	"fmt"
	"net/http"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Live"))
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// Get the username and password from the request body
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	// Check if the email and password are valid

	if email == "" {
		app.errorLog.Print("Missing email")
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}
	if password == "" {
		app.errorLog.Print("Missing password")
		http.Error(w, "Missing password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged in"))
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if firstName == "" {
		app.errorLog.Print("Missing first name")
		http.Error(w, "Missing first name", http.StatusBadRequest)
		return
	}
	if lastName == "" {
		app.errorLog.Print("Missing last name")
		http.Error(w, "Missing last name", http.StatusBadRequest)
		return
	}
	if email == "" {
		app.errorLog.Print("Missing email")
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}
	if password == "" {
		app.errorLog.Print("Missing password")
		http.Error(w, "Missing password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	uCreatedMsg := fmt.Sprintf("User created: %s %s", firstName, lastName)
	w.Write([]byte(uCreatedMsg))
}

func (app *application) userDetails(w http.ResponseWriter, r *http.Request) {
	app.errorLog.Print("Unauthorized")
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
