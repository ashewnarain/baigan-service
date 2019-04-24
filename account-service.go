package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
	for _, item := range accounts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Account{})
}

// CreateAccount creates an account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Printf("HTTP POST /accounts/%v\n", params["id"])
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.ID = params["id"]
	w.Header().Set("Content-Type", "application/json")
	id := CreateAccountRecord(account.ID, account.Firstname, account.Lastname, account.EmailAddress)
	json.NewEncoder(w).Encode(id)
}

// DeleteAccount deletes an account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Printf("HTTP DELETE /accounts/%v\n", params["id"])
	for index, item := range accounts {
		if item.ID == params["id"] {
			accounts = append(accounts[:index], accounts[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
