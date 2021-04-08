package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
	"fmt"
)

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)
	var port string = "8080"
//	var port string = os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")

	// css static routes
	router.Static("/css", "static/css")
	router.Static("/old/css", "static/css/old")
	// js static routes
	/*router.Static("/scripts", "static/scripts")
	router.Static("/media", "static/media")*/

	// serve classic themed pages
	router.GET("/old", func (c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/old/en")
	})
	router.GET("/old/en", func (c *gin.Context) {
		c.HTML(http.StatusOK, "en.index.html", nil)
	})
	router.GET("/old/jp", func (c *gin.Context) {
		c.HTML(http.StatusOK, "jp.index.html", nil)
	})
	
	router.POST("/old/submit", func (c *gin.Context) {
		// do something else plz
		var name string = c.PostForm("name")
		fmt.Println(name)
	})

	router.Run(":" + port)
}
