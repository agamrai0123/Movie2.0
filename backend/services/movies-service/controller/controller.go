package controllers

import (
	"net/http"

	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/models"
	"github.com/agamrai0123/Movie2.0/backend/services/movies-service/services"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	Service *services.MovieService
}

func NewMovieController() *MovieController {
	return &MovieController{
		Service: services.NewMovieService(),
	}
}

func (mc *MovieController) GetAllMovies(c *gin.Context) {
	movies, err := mc.Service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (mc *MovieController) GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := mc.Service.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := mc.Service.CreateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (mc *MovieController) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := mc.Service.UpdateMovie(id, &movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	if err := mc.Service.DeleteMovie(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
