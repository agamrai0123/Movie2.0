package models

type Movie struct {
	MovieID     string `json:"movie_id"`
	ScheduleID  string `json:"schedule_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Runtime     string `json:"runtime"`
}
