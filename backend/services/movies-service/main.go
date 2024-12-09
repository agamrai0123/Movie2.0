package main

import (
	"log"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/db"
	"github.com/agamrai0123/Movie2.0/backend/shared/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// log.Printf("Starting server on port %s", config.Port)
	// log.Printf("Starting server on port %s", config.DatabaseURL)
	// log.Printf("Starting server on port %s", config.Environment)
	database, err := db.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error opening db connection: %v", err)
	}
	defer db.CloseDB(database)

	// // List all movies
	// movies, err := listMovies(database)
	// if err != nil {
	// 	log.Fatalf("Error executing query: %v", err)
	// }
	// fmt.Println("All movies:")
	// for _, movie := range movies {
	// 	fmt.Printf("%s - %s (%s)\n", movie.MovieID, movie.Name, movie.Runtime)
	// }

	// router := gin.Default()

	// Load routes from the routes package
	// routes.RegisterMovieRoutes(router)

	// Start the server
	// port := "8080"
	// fmt.Printf("Starting server at port %s\n", port)
	// if err := router.Run(":" + port); err != nil {
	// 	log.Fatalf("Failed to start server: %s", err)
	// }
}
