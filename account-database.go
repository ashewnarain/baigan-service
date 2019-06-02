package main

import (
	"fmt"
	"time"
)

// CreateAccountRecord creates an account record
func CreateAccountRecord(accountID string, firstName string, lastName string, emailAddress string) string {
	id := "0"
	sqlStatement := `
		INSERT INTO smartbox.account (account_id, first_name, last_name, email_address)
		VALUES ($1, $2, $3, $4)
		RETURNING account_id`
	err := db.QueryRow(sqlStatement, accountID, firstName, lastName, emailAddress).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return id
}

// GetAccountRecords gets alls account records
func GetAccountRecords() []Account {
	rows, err := db.Query("SELECT * FROM smartbox.account LIMIT $1", 300)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var accounts []Account
	for rows.Next() {
		var accountID string
		var firstName string
		var lastName string
		var emailAddress string
		var createdTS time.Time
		var updatedTS time.Time

		err = rows.Scan(&accountID, &firstName, &lastName, &emailAddress, &createdTS, &updatedTS)
		if err != nil {
			// handle this error
			panic(err)
		}
		accounts = append(accounts,
			Account{
				ID:               accountID,
				Firstname:        firstName,
				Lastname:         lastName,
				EmailAddress:     emailAddress,
				CreatedTimestamp: createdTS,
				UpdatedTimestamp: updatedTS})
		fmt.Println(accountID, firstName, lastName, emailAddress, createdTS, updatedTS)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return accounts
}

// UpdateAccountRecord updates account record
func UpdateAccountRecord(accountID string, firstName string, lastName string, emailAddress string) string {
	id := "0"
	sqlStatement := `
		UPDATE smartbox.account
		SET first_name = $2, 
		last_name = $3,
		email_address = $4
		WHERE account_id = $1
		RETURNING account_id;`
	err := db.QueryRow(sqlStatement, accountID, firstName, lastName, emailAddress).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated record ID:", id)
	return id
}

// DeleteAccountRecord deletes an account record
func DeleteAccountRecord(accountID string) string {
	sqlStatement := `
		DELETE FROM smartbox.account
		WHERE account_id = $1
		RETURNING account_id;`
	var id string
	err := db.QueryRow(sqlStatement, accountID).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	return id
}
