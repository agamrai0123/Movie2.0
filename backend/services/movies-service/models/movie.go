package models

type Movie struct {
	MovieID     int    `json:"movie_id"`
	ScheduleID  int    `json:"schedule_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Runtime     string `json:"runtime"`
	// Genres      []string `json:"genres"`
}
