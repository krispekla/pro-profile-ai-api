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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing email"))
		return
	}
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing password"))
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing first name"))
		return
	}
	if lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing last name"))
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing email"))
		return
	}
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing password"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	uCreatedMsg := fmt.Sprintf("User created: %s %s", firstName, lastName)
	w.Write([]byte(uCreatedMsg))
}

func userDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}
