package repository

import (
	"database/sql"
	"fmt"
	"log"

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

func (r *MovieRepository) GetMovieByName(name string) (*models.Movie, error) {
	var movie *models.Movie
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies WHERE name = $1"
	row := r.DB.QueryRow(query, name)
	err := row.Scan(&movie.MovieID, &movie.ScheduleID, &movie.Name, &movie.Description, &movie.Runtime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepository) GetMovieByID(movieID uint) (*models.Movie, error) {
	var movie *models.Movie
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies WHERE id = $1"
	row := r.DB.QueryRow(query, movieID)
	err := row.Scan(&movie.MovieID, &movie.ScheduleID, &movie.Name, &movie.Description, &movie.Runtime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepository) CreateMovie(movie *models.Movie) error {
	query := "INSERT INTO movies (schedule_id, name, description, runtime) VALUES ($1, $2, $3, $4) RETURNING movie_id"
	err := r.DB.QueryRow(query, movie.ScheduleID, movie.Name, movie.Description, movie.Runtime).Scan(&movie.MovieID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// func UpdateMovie(db *sql.DB, movieID string, updatedMovie *models.Movie) error {

// }

func (r *MovieRepository) DeleteMovie(movieID uint) error {
	query := "DELETE FROM movies where id = $1"
	_, err := r.DB.Exec(query, movieID)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return err
	}
	return nil
}
