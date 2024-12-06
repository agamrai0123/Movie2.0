package models

type Movie struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageSrc    string `json:"image"`
}
