package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const keyLength = 8

func isInvalidKey(key string) bool {
	if len(key) > keyLength {
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

func insert(c *gin.Context) {
	redirectKey := uuid.New().String()
	redirectKey = strings.Replace(redirectKey, "-", "", -1)

	redirectKey = redirectKey[0:keyLength]

	if isInvalidKey(redirectKey) {
		c.JSON(500, "bad redirectKey")
		return
	}

	redirectURL, err := url.Parse(c.Query("redirectURL"))
	if err != nil {
		c.JSON(400, "bad redirectURL")
		return
	}

	if redirectURL.Host == "" {
		c.JSON(400, "bad redirectURL")
		return
	}

	err = insertIntoDB(redirectKey, redirectURL.Host)
	if err != nil {
		c.JSON(500, fmt.Sprintf("could not insert into db: %s", err.Error()))
		return
	}

	c.JSON(200, fmt.Sprintf("Inserted %s!", redirectKey))
}

func home(c *gin.Context) {
	c.JSON(200, "Hello world!")
}

func redirect(c *gin.Context) {
	redirectKey := c.Param("redirect")
	fmt.Println("Key: ", redirectKey)

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

	fmt.Println("Redirect URL: ", redirectURL)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s", redirectURL))
}

func allRedirects(c *gin.Context) {
	redirects, err := getAllRedirects()
	if err != nil {
		c.JSON(400, fmt.Sprintf("err was %s", err))
		return
	}

	c.JSON(200, redirects)
}

func handleDelete(c *gin.Context) {
	redirectKey := c.Param("redirectKey")
	if isInvalidKey(redirectKey) {
		c.JSON(400, "Bad redirectKey")
		return
	}

	err := deleteRedirect(redirectKey)
	if err != nil {
		c.JSON(400, "Unable to delete redirect key")
		return
	}

	c.JSON(200, "OK")
}

func handleUpdate(c *gin.Context) {
	redirectKey := c.Query("redirectKey")
	fmt.Println(redirectKey)
	if isInvalidKey(redirectKey) {
		c.JSON(400, "Bad redirectKey")
		return
	}

	// oldRedirectURL, err := url.Parse(c.Query("oldRedirectURL"))
	// if err != nil {
	// 	c.JSON(400, "bad old redirectURL")
	// 	return
	// }

	// if oldRedirectURL.Host == "" {
	// 	c.JSON(400, "bad old redirectURL")
	// 	return
	// }

	newRedirectURL, err := url.Parse(c.Query("newRedirectURL"))
	if err != nil {
		c.JSON(400, "bad new redirectURL")
		return
	}

	if newRedirectURL.Host == "" {
		c.JSON(400, "the url cannot be empty")
		return
	}

	err = updateRedirect(redirectKey, newRedirectURL.Host)
	if err != nil {
		c.JSON(400, fmt.Sprintf("error updating redirect %s", err.Error()))
		return
	}

	c.JSON(200, "OK")
}

func handleSeed(c *gin.Context) {
	err := seedDB()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "Seeded")
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/:redirect", redirect)
	router.GET("/home", home)

	routerAdmin := gin.Default()
	routerAdmin.Use(cors.Default())

	routerAdmin.GET("/insert", insert)
	routerAdmin.POST("/insert", insert)
	routerAdmin.DELETE("/:redirectKey", handleDelete)
	routerAdmin.GET("/all", allRedirects)
	routerAdmin.PUT("/update", handleUpdate)
	routerAdmin.PUT("/seed", handleSeed)

	// seperate out the admin routes from the normal routes
	go func() {
		_ = routerAdmin.Run(":5001")
	}()

	_ = router.Run(":5000")
}
