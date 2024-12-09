package repository

import (
	"database/sql"
	"fmt"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"

	_ "github.com/lib/pq"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository() *MovieRepository {
	connStr := "user=postgres password=yourpassword dbname=movies_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	return &MovieRepository{
		DB: db,
	}
}

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
	query := "SELECT * FROM movies"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.MovieID, &movie.Name, &movie.Runtime, &movie.ScheduleID, &movie.Description); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
