package services

import (
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/repository"
)

type MovieService struct {
	Repo *repository.MovieRepository
}

func NewMovieService() *MovieService {
	return &MovieService{
		Repo: repository.NewMovieRepository(),
	}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.Repo.GetAllMovies()
}

// func (s *MovieService) GetMovieByID(id string) (*models.Movie, error) {
// 	return s.Repo.GetMovieByID(id)
// }
