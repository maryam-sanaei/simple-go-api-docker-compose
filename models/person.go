package models

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
