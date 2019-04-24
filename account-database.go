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
