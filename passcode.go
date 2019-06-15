package main

import "time"

// Passcode passcode for box
type Passcode struct {
	ProductID       string    `json:"product_id,omitempty"`
	PassCode        string    `json:"pass_code,omitempty"`
	InsertID        string    `json:"insert_id,omitempty"`
	InsertTimestamp time.Time `json:"insert_timestamp,omitempty"`
}
