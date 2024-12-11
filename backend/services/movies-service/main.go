package main

import (
	"fmt"
	"log"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.RegisterMovieRoutes(router)

	port := "8080"
	fmt.Printf("Starting server at port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
