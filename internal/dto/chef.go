package dto

import "time"

type Chef struct {
	Name        string    `json:"name"`
	Catchphrase string    `json:"catchphrase"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"` // domainに関係ないフィールドもdtoで表現できる
}

func NewChef() *Chef {
	return &Chef{}
}
