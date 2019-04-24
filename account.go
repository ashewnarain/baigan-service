package main

import "time"

// Account account details
type Account struct {
	ID               string    `json:"id,omitempty"`
	Firstname        string    `json:"first_name,omitempty"`
	Lastname         string    `json:"last_name,omitempty"`
	EmailAddress     string    `json:"email_address,omitempty"`
	CreatedTimestamp time.Time `json:"created_timestamp,omitempty"`
	UpdatedTimestamp time.Time `json:"updated_timestamp,omitempty"`
}
