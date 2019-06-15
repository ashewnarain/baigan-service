package main

import "time"

// Account account details
type Account struct {
	EmailAddress     string    `json:"email_address,omitempty"`
	Firstname        string    `json:"first_name,omitempty"`
	Lastname         string    `json:"last_name,omitempty"`
	CreatedTimestamp time.Time `json:"created_timestamp,omitempty"`
	UpdatedTimestamp time.Time `json:"updated_timestamp,omitempty"`
}
