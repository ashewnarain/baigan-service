package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"io/ioutil"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	port = "80"
)

var logger *log.Logger
var accounts []Account

func main() {
	InitializeLogging()

	// temporary test data
	accounts = append(accounts, Account{ID: "1", Firstname: "Anthony", Lastname: "Shewnarain", EmailAddress: "anthony.shewnarain@gmail.com"})

	router := mux.NewRouter()
	router.HandleFunc("/", SayHello).Methods("GET")
	InitializeAccountEndpoints(router)

	// start the server
	done := make(chan bool)
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, router))
	}()
	logger.Printf("Baigan Service started at port %v...", port)
	<-done
}

// InitializeLogging initialize the logger
func InitializeLogging() {
	directory := "logs"
	fileName := "baigan-service.log"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile(directory+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(f, "", log.Ldate|log.Ltime)
	logger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   directory + "/" + fileName,
		MaxSize:    10, // megabytes
		MaxBackups: 1,
		MaxAge:     14,   //days
		Compress:   true, // disabled by default
	}))
}

// InitializeAccountEndpoints initialize accounts resources
func InitializeAccountEndpoints(r *mux.Router) {
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", GetAccount).Methods("GET")
	r.HandleFunc("/accounts/{id}", CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", DeleteAccount).Methods("DELETE")
	return
}

// SayHello Simple ping
func SayHello(w http.ResponseWriter, r *http.Request) {
	logger.Println("HTTP GET /")
	b, err := ioutil.ReadFile("README.md") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	readMe := string(b)
	fmt.Fprintln(w, readMe)
}

// GetAccounts gets list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	logger.Println("HTTP GET /accounts")
	w.Header().Set("Content-Type", "application/json")
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
	accounts = append(accounts, account)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
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

// Account account details
type Account struct {
	ID           string `json:"id,omitempty"`
	Firstname    string `json:"first_name,omitempty"`
	Lastname     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}
