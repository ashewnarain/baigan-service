package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	db "github.com/ashewnarain/baigan-service/database"
)

// CreateAccountRecord creates an account record
func CreateAccountRecord(emailAddress string, firstName string, lastName string) (string, error) {
	query := `
		INSERT INTO smartbox.account (email_address, first_name, last_name)
		VALUES ($1, $2, $3)
		RETURNING email_address`
	stmt, err := db.SQLDB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var message string
	err = stmt.QueryRow(emailAddress, firstName, lastName).Scan(&message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("New record ID is:", emailAddress)
	return message, nil
}

// GetAccountRecords gets alls account records
func GetAccountRecords() []Account {
	rows, err := db.SQLDB.Query("SELECT * FROM smartbox.account LIMIT $1", 300)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var accounts []Account
	for rows.Next() {
		var emailAddress string
		var firstName string
		var lastName string
		var createdTS time.Time
		var updatedTS time.Time

		err = rows.Scan(&emailAddress, &firstName, &lastName, &createdTS, &updatedTS)
		if err != nil {
			// handle this error
			panic(err)
		}
		accounts = append(accounts,
			Account{
				EmailAddress:     emailAddress,
				Firstname:        firstName,
				Lastname:         lastName,
				CreatedTimestamp: createdTS,
				UpdatedTimestamp: updatedTS})
		fmt.Println(emailAddress, firstName, lastName, createdTS, updatedTS)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return accounts
}

//GetAccountRecord returns single account
func GetAccountRecord(id string) (Account, error) {
	sqlStatement := `
		SELECT email_address, first_name, last_name, created_ts, updated_ts 
		FROM smartbox.account 
		WHERE email_address=$1;`
	row := db.SQLDB.QueryRow(sqlStatement, id)
	var emailAddress string
	var firstName string
	var lastName string
	var createdTS time.Time
	var updatedTS time.Time
	switch err := row.Scan(&emailAddress, &firstName, &lastName, &createdTS, &updatedTS); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Account{}, err
	case nil:
		return Account{
			EmailAddress:     emailAddress,
			Firstname:        firstName,
			Lastname:         lastName,
			CreatedTimestamp: createdTS,
			UpdatedTimestamp: updatedTS}, nil
	default:
		return Account{}, err
	}
}

// UpdateAccountRecord updates account record
func UpdateAccountRecord(account Account) (string, error) {
	var emailAddress string
	sqlStatement := `
		UPDATE smartbox.account
		SET first_name = $2, 
		last_name = $3
		WHERE email_address = $1
		RETURNING email_address;`
	fmt.Println(account.EmailAddress)
	err := db.SQLDB.QueryRow(sqlStatement, account.EmailAddress, account.Firstname, account.Lastname).Scan(&emailAddress)
	if err != nil {
		fmt.Println("No rows were updated!")
		return emailAddress, err
	}
	fmt.Println("Updated record ID:", emailAddress)
	return emailAddress, nil
}

// DeleteAccountRecord deletes an account record
func DeleteAccountRecord(emailAddress string) string {
	sqlStatement := `
		DELETE FROM smartbox.account
		WHERE email_address = $1
		RETURNING 'SUCCESS';`
	var message string
	err := db.SQLDB.QueryRow(sqlStatement, emailAddress).Scan(&message)
	if err != nil {
		panic(err)
	}
	fmt.Println(message)
	return message
}
