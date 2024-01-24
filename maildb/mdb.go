/*
Author: Adrian Potra
Version 1.0

Note: we will import sql light driver from github - github.com/mattn/go-sqlite3
*/

package mdb

import (
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

// creating email entry structure - will be utilized for both reading and adding to the db

type EmailEntry struct {
	Id          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}

// function to create db - where we're going to execute the queries to create the db tables

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
	  CREATE TABLE emails (
          id  INTEGER PRIMARY KEY,
		  email TEXT UNIQUE,
		  confirmed_at INTEGER,
		  opt_out 	INTEGER
	  );
	`)
	if err != nil { // there can be multiple errors that could be checked but we are only checking if the database exists already
		if sqlError, ok := err.(sqlite3.Error); ok { // we take the error and we cast it into a sqlite type. once we have the error we check that it's not equal to 1 which means db exists - if we don't have 1, it means that something else happened and we log the err
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else { // if the table exists, we need to handle all the errors
			log.Fatal(err)
		}
	}
}

// function that creates an email entry struct from db row

func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var id int64
	var email string
	var confirmedAt int64
	var optOut bool
	// function to read out of the db - we use pointers to read out of the vars declared above
	err := row.Scan(&id, &email, &confirmedAt, &optOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// convert the time into an appropriate time struct - unix time
	t := time.Unix(confirmedAt, 0)
	return &EmailEntry{Id: id, Email: email, ConfirmedAt: &t, OptOut: optOut}, nil

}

// creating DB CRUD operations

func CreateEmail(db *sql.DB, email string) error {
	// execute the query and insert email data - the email object will be substituted for the '?' in query, we set 0 confirmed at time to indicate email has not been confirmed and default opt out is 0
	// and id - it will be set by sqlite
	_, err := db.Exec(`
     INSERT INTO 
	 emails(email, confirmed_at, opt_out)
	 VALUES(?, 0, false)`, email)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil // indicates that query was ok
}

// read email func
func GetEmail(db *sql.DB, email string) (*EmailEntry, error) {
	rows, err := db.Query(`
    SELECT id, email, confirmed_at, opt_out
	FROM emails
	WHERE email = ?`, email)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close() // here we need to close the db rows because the db.Query functions leaves it open to read the rows, so we have to close gracefully after
	// iterate through the wors to retrieve the emails
	for rows.Next() {
		return emailEntryFromRow(rows)
	}
	return nil, nil // if both return values are nil, it indicates that there was no email or error
}

// update email - upsert operation
func UpdateEmail(db *sql.DB, entry EmailEntry) error {
	// get time as unix time
	t := entry.ConfirmedAt.Unix()
	_, err := db.Exec(`
     INSERT INTO 
	 emails(email, confirmed_at, opt_out)
	 VALUES(?, ?, ?)
	 ON CONFLICT(email) DO UPDATE SET
	 confirmed_at = ?,
	 opt_out = ?`, entry.Email, t, entry.OptOut, t, entry.OptOut)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

// delete email - for this app we won't delete email from db
// but we just set the opt out value to be true, since it's a mailing list
func DeleteEmail(db *sql.DB, email string) error {

	_, err := db.Exec(`
	UPDATE emails
	SET opt_out=true
	WHERE email=?`, email)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// batch query parameters

type GetEmailBatchQueryParams struct {
	Page  int // page nr - for pagination
	Count int // nr of emails that are supposed to be returned
}

func GetEmailBatch(db *sql.DB, params GetEmailBatchQueryParams) ([]EmailEntry, error) {
	// create empty slice
	var empty []EmailEntry
	// query the db
	rows, err := db.Query(`
   SELECT id, email, confirmed_at, opt_out
   FROM emails
   WHERE opt_out = false
   ORDER BY id ASC
   LIMIT ? OFFSET ?`, params.Count, (params.Page-1)*params.Count) // we will offset by the page that we're on - index starts at 0

	if err != nil {
		log.Println(err)
		return empty, err
	}
	defer rows.Close() // here we need to close the db rows because the db.Query functions leaves it open to read the rows, so we have to close gracefully after
	emails := make([]EmailEntry, 0, params.Count)
	// we iterate through each email
	for rows.Next() {
		email, err := emailEntryFromRow(rows)
		if err != nil {
			return nil, err
		}
		emails = append(emails, *email)
	}

	return emails, nil
}
