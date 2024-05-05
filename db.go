package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func getUrlFromKey(redirectKey string) (string, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/urls")

	if err != nil {
		return "", err
	}

	defer db.Close()

	rows := db.QueryRow(`select redirectURL from links where redirectKey = ?`, redirectKey)

	redirectURL := ""
	if err := rows.Scan(&redirectURL); err != nil {
		return "", err
	}

	return redirectURL, nil
}

func insertIntoDB(redirectKey, redirectURL string) error {
	if isInvalidKey(redirectKey) {
		return errors.New("Invalid redirect key!")
	}

	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/urls")
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Query(`insert into links (redirectKey, redirectURL) values (?, ?)`, redirectKey, redirectURL)
	if err != nil {
		return err
	}

	return nil
}
