package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const hostname = "mysqldb"
const dbname = "mydb"
const dbuser = "myuser"
const dbpass = "mypass"

func GetPersons() []Person {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+hostname+":3306)/"+dbname)

	// Connection issue, so handle it
	if err != nil {

		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	// Close till after this function has finished
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	persons := []Person{}

	for rows.Next() {

		var person Person

		// Scan every row and put it to the Person intance, it must be match ordered with the person table fields
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)

		if err != nil {
			panic(err.Error())
		}

		// Append new instance to persons slice
		persons = append(persons, person)

	}

	return persons

}

func AddPerson(person Person) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+hostname+":3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// Close till after this function has finished
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO person (id,firstname,lastname) VALUES (NULL,?,?)", person.FirstName, person.LastName)

	// If there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
