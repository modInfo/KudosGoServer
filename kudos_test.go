package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"testing"
)

// TODO:
// add setUp & tearDown methods for the test DB

// var db sql.SQL

func TestReverse(t *testing.T) {
	setUp()

	testButton := KudoButton{
		URL: "http://coucou/",
		ID:  "test-btn",
	}

	if getCurrentCount(testButton.ID) != 0 {
		t.Errorf("Default value for an inexistent button should be 0 kudo")
	}

	// Test default value for a button
	testButton.create()
	if testButton.KudoCount != 0 || getCurrentCount(testButton.ID) != 0 {
		t.Errorf("Default value for a button should be 0 kudo")
	}

	// Test increase for inexistent button, should do nothing and return 0
	affectedButton := increaseKudoButton("inexistent-btn")
	if affectedButton != 0 || getCurrentCount(testButton.ID) != 0 {
		t.Errorf("Increasing an inexistent button should do nothing")
	}

	// Test increase for existent button
	affectedButton = increaseKudoButton(testButton.ID)
	if affectedButton != 1 || getCurrentCount(testButton.ID) != 1 {
		t.Errorf("Increasing a button should increment the kudoCount")
	}

}

func setUp() {
	databaseUrl = "file:test.db?cache=shared&mode=memory"

	db, err := sql.Open("sqlite3", databaseUrl)

	body, err := ioutil.ReadFile("initdb.sql")
	checkErr(err)

	db.Exec(string(body))
}

func tearDown() {
	// Close the DB connection
	// db.Close()
}