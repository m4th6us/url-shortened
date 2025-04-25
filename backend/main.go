package main

import (
	"urlshorten/database"
	"urlshorten/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMongo()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 		[]string{"http://localhost:8081"},
		AllowMethods: 		[]string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: 		[]string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: 	true,
	}))

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.Redirect)

	r.Run(":8080")
}