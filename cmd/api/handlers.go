package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Live"))
}

func login(w http.ResponseWriter, r *http.Request) {
	// Get the username and password from the request body
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	// Check if the email and password are valid

	if email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}
	if password == "" {
		http.Error(w, "Missing password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged in"))
}

func register(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("first_name")
	lastName := r.PostFormValue("last_name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if firstName == "" {
		http.Error(w, "Missing first name", http.StatusBadRequest)
		return
	}
	if lastName == "" {
		http.Error(w, "Missing last name", http.StatusBadRequest)
		return
	}
	if email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}
	if password == "" {
		http.Error(w, "Missing password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	uCreatedMsg := fmt.Sprintf("User created: %s %s", firstName, lastName)
	w.Write([]byte(uCreatedMsg))
}

func userDetails(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
