package models

import "time"

type Movie struct {
	MovieID     uint      `json:"movie_id"`
	ScheduleID  uint      `json:"schedule_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Runtime     time.Time `json:"runtime"`
	// Genres      []string `json:"genres"`
}
