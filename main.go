package main

import (
	_ "go-gin-api/docs"
	"go-gin-api/repository"
	"go-gin-api/router"
	"log"

	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load()

	if err := repository.InitDB(); err != nil {
		log.Fatal("failed to connect database:", err)
	}

	r := router.SetRouter()

	r.Run(":8080")
}
