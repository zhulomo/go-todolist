package main

import (
	_ "go-gin-api/docs"
	"go-gin-api/repository"
	"go-gin-api/router"
	"log"
)

func main() {

	if err := repository.InitDB(); err != nil {
		log.Fatal("failed to connect database:", err)
	}

	r := router.SetRouter()

	r.Run(":8080")
}
