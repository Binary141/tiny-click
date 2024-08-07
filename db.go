package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type redirects struct {
	Key   string `json:"redirectKey"`
	Value string `json:"redirectValue"`
}

const dbURL = "root:secret@tcp(db:3306)/urls"

func getUrlFromKey(redirectKey string) (string, error) {
	db, err := sql.Open("mysql", dbURL)

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

	db, err := sql.Open("mysql", dbURL)
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

func getAllRedirects() ([]redirects, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return []redirects{}, err
	}

	defer db.Close()
	rows, err := db.Query(`select redirectKey, redirectURL from links`)
	if err != nil {
		return []redirects{}, nil
	}

	var allRedirects []redirects

	for rows.Next() {
		tmp := redirects{}

		if err := rows.Scan(&tmp.Key, &tmp.Value); err != nil {
			return allRedirects, err
		}

		allRedirects = append(allRedirects, tmp)
	}

	return allRedirects, nil
}

func deleteRedirect(key string) error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec(`delete from links where redirectKey = ?`, key)
	if err != nil {
		return err
	}

	return nil
}

func updateRedirect(redirectKey, redirectURL string) error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec(`update links set redirectURL = ? where redirectKey = ?`, redirectURL, redirectKey)
	if err != nil {
		return err
	}

	return nil
}

func seedDB() error {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec(`
		create table if not exists links (
			id int primary key auto_increment,
			redirectURL varchar(255),
			redirectKey varchar(128) unique
		)`)
	if err != nil {
		return err
	}

	return nil
}
