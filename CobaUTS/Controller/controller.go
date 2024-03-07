package Controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "CobaUTS/Model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username == "" || password == "" {
		log.Println("Error: Email and Password Missing")
		http.Error(w, "Bad Request: Email and Password Missing", http.StatusBadRequest)
		return
	}

	var (
		dbPassword string
	)

	query := "SELECT passwords FROM users WHERE email = ?"
	err := db.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "Can't find username")
			return
		}
		sendErrorResponse(w, "Internal Server Error")
		return
	}

	if password != dbPassword {
		sendErrorResponse(w, "Wrong Password")
		return
	}

	platform := r.Header.Get("platform")
	fmt.Fprintf(w, "Success login from %s", platform)
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.SuccessResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
