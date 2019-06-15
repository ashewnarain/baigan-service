package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	db "github.com/ashewnarain/baigan-service/database"
)

// CreatePasscodeRecord creates a passcode record
func CreatePasscodeRecord(productID string, passcode string, insertID string) (string, error) {
	query := `
		INSERT INTO smartbox.pass_code (product_id, pass_code, insert_id)
		VALUES ($1, $2, $3)
		RETURNING insert_ts`
	stmt, err := db.SQLDB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var message time.Time
	err = stmt.QueryRow(productID, passcode, insertID).Scan(&message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("New record ID is:", message)
	return message.String(), nil
}

//GetPasscodeRecord returns last passcode for product
func GetPasscodeRecord(id string) (Passcode, error) {
	sqlStatement := `
		SELECT product_id, pass_code, insert_id, insert_ts
		FROM smartbox.pass_code 
		WHERE product_id=$1
		ORDER BY insert_ts desc
		LIMIT 1;`
	row := db.SQLDB.QueryRow(sqlStatement, id)
	var productID string
	var passcode string
	var insertID string
	var insertTS time.Time
	switch err := row.Scan(&productID, &passcode, &insertID, &insertTS); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Passcode{}, err
	case nil:
		return Passcode{
			ProductID:       productID,
			PassCode:        passcode,
			InsertID:        insertID,
			InsertTimestamp: insertTS}, nil
	default:
		return Passcode{}, err
	}
}
