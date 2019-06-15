package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreatePasscode creates a passcode
func CreatePasscode(w http.ResponseWriter, r *http.Request) {
	logger.Printf("HTTP POST /passcodes")
	var passcode Passcode
	err := json.NewDecoder(r.Body).Decode(&passcode)
	// error decoding body
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	id, err := CreatePasscodeRecord(passcode.ProductID, passcode.PassCode, passcode.InsertID)
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

// GetPasscode gets the last passcode
func GetPasscode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logger.Printf("HTTP GET /passcodes/%v\n", params["id"])
	id := params["id"]
	account, err := GetPasscodeRecord(id)
	if err != nil {
		logger.Printf(err.Error())
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			// error calling database
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(account)
}
