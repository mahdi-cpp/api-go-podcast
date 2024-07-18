package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	podcast "github.com/mahdi-cpp/api-go-podcast/api"
)

var (
	router = gin.Default()
)

// Run will Go-Instagram the server
func Run() {
	router.Use(CORSMiddleware())
	getRoutes()
	router.Run(":8096")
}

// getRoutes will create our api of our entire application
// this way every group of api can be defined in their own file
// so this one won't be so messy

func getRoutes() {

	v1 := router.Group("/v1")

	podcast.AddMusicRoutes(v1)
	podcast.AddPodcastRoutes(v1)
	//javascript.AddJavascriptRoutes(v1)
	//addGalleryRoutes(v1)

	//v2 := router.Group("/v2")
	//addStoryRoutes(v2)
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
