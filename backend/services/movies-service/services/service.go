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

func (s *MovieService) GetMovieByName(name string) (*models.Movie, error) {
	return s.Repo.GetMovieByName(name)
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
	return s.Repo.GetMovieByID(id)
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.Repo.CreateMovie(movie)
}

// func (s *MovieService) UpdateMovie(id uint, updatedMovie *models.Movie) error {
// 	return s.Repo.UpdateMovie(id, updatedMovie)
// }

func (s *MovieService) DeleteMovie(id uint) error {
	return s.Repo.DeleteMovie(id)
}
