// main.go

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set the router as the default one provided by Gin
	router := gin.Default()
	
	// using LoadHTMLGlob instead of LoadHTMLFiles allows us to use a string pattern in the parameters instead of raw string
	router.LoadHTMLGlob("templates/*")

	// inline GET route for the landing page which server the index.html template with necessary data - `title`
	router.GET("/", func(c *gin.Context){
		articles := getAllArticles()
		println(articles)
	
		// call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// use the index.html template
			"index.html",
			// Pass the data the page requires
			gin.H{
				"title": "Home Page",
				"payload": articles,
			})
	})

	// initializeRoutes()

	// start serving the application
	router.Run("0.0.0.0:8080")
}