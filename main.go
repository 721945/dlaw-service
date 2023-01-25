package main

import (
	"log"
	"net/http"

	"github.com/721945/dlaw-service/configs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// database_configs
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	db, err := configs.InitialDatbase()

	if err != nil {
		log.Fatal(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
