package models

import (
	"time"
)

type Episode struct {
	Id                int       `db:"id"`
	StoryNumber       string    `db:"story_number"`
	SeriesNumber      int       `db:"series_number"`
	EpisodeNumber     int       `db:"episode_number"`
	Title             string    `db:"title"`
	DoctorActor       int       `db:"doctor_actor"`
	DirectedBy        int       `db:"directed_by"`
	WrittenBy         int       `db:"written_by"`
	OriginalAirDate   time.Time `db:"original_air_date"`
	Viewers           float64   `db:"viewers"`
	AppreciationIndex int       `db:"appreciation_index"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type Person struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
