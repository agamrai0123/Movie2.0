package routes

import (
	controllers "github.com/agamrai0123/Movie2.0/backend/services/movies-service/controller"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine) {
	// Create a new instance of MovieController
	movieController := controllers.NewMovieController()

	// Define endpoints and associate them with controller methods
	router.GET("/movies", movieController.GetAllMovies)       // GET /movies
	router.GET("/movies/:id", movieController.GetMovieByID)   // GET /movies/:id
	router.POST("/movies", movieController.CreateMovie)       // POST /movies
	router.PUT("/movies/:id", movieController.UpdateMovie)    // PUT /movies/:id
	router.DELETE("/movies/:id", movieController.DeleteMovie) // DELETE /movies/:id
}
