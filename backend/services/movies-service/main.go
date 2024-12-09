package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/db"
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/routes"
	"github.com/agamrai0123/Movie2.0/backend/shared/config"
	"github.com/gin-gonic/gin"
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

	// activeConnections, err := getActiveConnections(database)
	// if err != nil {
	// 	log.Fatalf("Error executing query: %v", err)
	// }
	// fmt.Printf("Number of active connections: %d\n", activeConnections)

	// // List all movies
	// movies, err := listMovies(database)
	// if err != nil {
	// 	log.Fatalf("Error executing query: %v", err)
	// }
	// fmt.Println("All movies:")
	// for _, movie := range movies {
	// 	fmt.Printf("%s - %s (%s)\n", movie.MovieID, movie.Name, movie.Runtime)
	// }

	router := gin.Default()

	// Load routes from the routes package
	routes.RegisterMovieRoutes(router)

	// Start the server
	port := "8080"
	fmt.Printf("Starting server at port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func listMovies(db *sql.DB) ([]*models.Movie, error) {
	var movies []*models.Movie

	// Use db.Query instead of db.QueryRow to handle multiple rows
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close() // Ensure the rows are closed to avoid resource leaks

	// Loop over the result set and populate the movies slice
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.MovieID, &movie.ScheduleID, &movie.Name, &movie.Description, &movie.Runtime)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		movies = append(movies, &movie)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return movies, nil
}

// func getActiveConnections(db *sql.DB) (int, error) {
// 	var activeConnections int
// 	query := "SELECT COUNT(*) FROM pg_stat_activity WHERE state = 'active'"
// 	err := db.QueryRow(query).Scan(&activeConnections)
// 	if err != nil {
// 		log.Fatalf("Error executing query: %v", err)
// 		return -1, err
// 	}
// 	return activeConnections, nil
// }
