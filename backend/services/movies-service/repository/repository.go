package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/database"
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"
	"github.com/agamrai0123/Movie2.0/backend/shared/config"

	_ "github.com/lib/pq"
)

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository() *MovieRepository {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := database.InitDB(config.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	n, err := database.GetActiveConnections(db)
	if err != nil {
		log.Fatalf("Error getting active connections: %v", err)
	}
	log.Printf("Active connections: %d", n)

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
		if err := rows.Scan(&movie.MovieID, &movie.ScheduleID, &movie.Name, &movie.Description, &movie.Runtime); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *MovieRepository) GetMovieByName(name string) (*models.Movie, error) {
	movie := &models.Movie{}
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

func (r *MovieRepository) GetMovieByID(movieID int) (*models.Movie, error) {
	movie := &models.Movie{}
	query := "SELECT movie_id, schedule_id, name, description, runtime FROM movies WHERE movie_id = $1"
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

func (r *MovieRepository) UpdateMovie(movieID int, updatedMovie *models.Movie) error {
	query := "UPDATE movies SET schedule_id =$1 name = $2, description = $3, runtime = $4 WHERE movie_id = $5"
	_, err := r.DB.Exec(query, updatedMovie.ScheduleID, updatedMovie.Name, updatedMovie.Description, updatedMovie.Runtime, movieID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

func (r *MovieRepository) DeleteMovie(movieID int) error {
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
