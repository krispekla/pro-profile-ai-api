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

		if email == "" || password == "" {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		msg := fmt.Sprintf("User %s is logged in", email)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	}
}

func register(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		if firstName == "" || lastName == "" || email == "" || password == "" {
			app.ClientError(w, http.StatusBadRequest)
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
