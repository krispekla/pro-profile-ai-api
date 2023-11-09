package main

import (
	"fmt"
	"net/http"

	"github.com/krispekla/pro-profile-ai-api/config"
)

func ping(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Live"))
	}
}

func login(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the username and password from the request body
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		// Check if the email and password are valid

		if email == "" {
			app.ErrorLog.Print("Missing email")
			http.Error(w, "Missing email", http.StatusBadRequest)
			return
		}
		if password == "" {
			app.ErrorLog.Print("Missing password")
			http.Error(w, "Missing password", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged in"))
	}
}

func register(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		if firstName == "" {
			app.ErrorLog.Print("Missing first name")
			http.Error(w, "Missing first name", http.StatusBadRequest)
			return
		}
		if lastName == "" {
			app.ErrorLog.Print("Missing last name")
			http.Error(w, "Missing last name", http.StatusBadRequest)
			return
		}
		if email == "" {
			app.ErrorLog.Print("Missing email")
			http.Error(w, "Missing email", http.StatusBadRequest)
			return
		}
		if password == "" {
			app.ErrorLog.Print("Missing password")
			http.Error(w, "Missing password", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		uCreatedMsg := fmt.Sprintf("User created: %s %s", firstName, lastName)
		w.Write([]byte(uCreatedMsg))
	}
}

func userDetails(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.ErrorLog.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
