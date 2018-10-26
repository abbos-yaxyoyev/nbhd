package models

type Task struct {
	Id            string
	Title         string
	Category      int
	Location      []float64
	Description   string
	Time          int
	Creator       string
	Performer     string
	Encouragement int
	Pay           float64
	Created       string
	Archived      bool
}
