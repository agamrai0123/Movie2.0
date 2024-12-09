package db_store

import (
	"database/sql"
	"log"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"
)

func ListAllMovies(db *sql.DB) ([]*models.Movie, error) {
	var movies []*models.Movie

	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

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

func GetMovieByName(db *sql.DB, name string) (*models.Movie, error) {
	var movies *models.Movie
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies WHERE name = $1"
	row := db.QueryRow(query, name)
	err := row.Scan(&movies.MovieID, &movies.ScheduleID, &movies.Name, &movies.Description, &movies.Runtime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}
	return movies, nil
}

func GetMovieByID(db *sql.DB, movieID string) (*models.Movie, error) {
	var movies *models.Movie
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies WHERE id = $1"
	row := db.QueryRow(query, movieID)
	err := row.Scan(&movies.MovieID, &movies.ScheduleID, &movies.Name, &movies.Description, &movies.Runtime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}
	return movies, nil
}

func CreateMovie(db *sql.DB, movie *models.Movie) error {
	query := "INSERT INTO movies (schedule_id, name, description, runtime) VALUES ($1, $2, $3, $4) RETURNING movie_id"
	err := db.QueryRow(query, movie.ScheduleID, movie.Name, movie.Description, movie.Runtime).Scan(&movie.MovieID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// func UpdateMovie(db *sql.DB, movieID string, updatedMovie *models.Movie) error {

// }

// func DeleteMovie(db *sql.DB, movieID string) error
