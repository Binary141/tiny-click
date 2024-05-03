package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func insert(c *gin.Context) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/urls")

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	defer db.Close()

	_, err = db.Query(`insert into links (test) values (10)`)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "hello")
}

func isInvalidKey(key string) bool {
	if len(key) > 8 {
		return true
	}

	// make sure the string only has alpha numerics
	match, err := regexp.MatchString("^[0-9a-zA-Z]+$", key)
	if err != nil {
		return true
	}

	// if the regex matched, then it was a valid string.
	// we want to see if it was invalid
	return !match
}

func home(c *gin.Context) {
	c.JSON(200, "Hello world!")
}

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

func redirect(c *gin.Context) {
	redirectKey := c.Param("redirect")

	if isInvalidKey(redirectKey) {
		fmt.Println("Key was invalid!")
		redirectHome := fmt.Sprintf("http://%s%s", c.Request.Host, "/home")
		c.Redirect(http.StatusMovedPermanently, redirectHome)
		return
	}

	redirectURL, err := getUrlFromKey(redirectKey)
	if err != nil {
		fmt.Printf("Error was %s\n", err.Error())
		redirectHome := fmt.Sprintf("http://%s%s", c.Request.Host, "/home")
		c.Redirect(http.StatusMovedPermanently, redirectHome)
		return
	}

	c.Redirect(http.StatusMovedPermanently, redirectURL)
}

func main() {
	router := gin.Default()
	router.GET("/home", home)
	router.GET("/insert", insert)
	router.GET("/:redirect", redirect)

	_ = router.Run(":5000")
}
