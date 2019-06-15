package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateAccount creates an account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	logger.Printf("HTTP POST /accounts")
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	// error decoding body
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	id, err := CreateAccountRecord(account.EmailAddress, account.Firstname, account.Lastname)
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

// UpdateAccount update an account
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	logger.Printf("HTTP PUT /accounts")
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	// error decoding body
	if err != nil {
		logger.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	id, err := UpdateAccountRecord(account)
	w.Header().Set("Content-Type", "application/json")
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
	json.NewEncoder(w).Encode(&id)
}

// GetAccounts gets list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	logger.Println("HTTP GET /accounts")
	w.Header().Set("Content-Type", "application/json")
	accounts := GetAccountRecords()
	json.NewEncoder(w).Encode(accounts)
}

// GetAccount gets a single account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logger.Printf("HTTP GET /accounts/%v\n", params["id"])
	id := params["id"]
	account, err := GetAccountRecord(id)
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

// DeleteAccount deletes an account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Printf("HTTP DELETE /accounts/%v\n", params["id"])
	id := DeleteAccountRecord(params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}
