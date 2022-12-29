package main

import (
	"backend/models"
	"encoding/json"
	"net/http"
	"time"
)

type Payload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	payload = Payload{
		Status:  "active",
		Message: "Go movies list",
		Version: "1.0",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	rd, _ := time.Parse("2006-01-02", "2000-09-22")
	harryPotter := models.Movie{
		ID:          1,
		Title:       "Harry Potter and Sorcerers Stone",
		Description: "An Awesome Movie",
		RunTime:     150,
		MPAARating:  "PG-13",
		ReleaseDate: rd,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	rd, _ = time.Parse("2006-01-02", "1994-05-28")
	terminator := models.Movie{
		ID:          2,
		Title:       "Terminator",
		Description: "An Awesome Movie",
		RunTime:     130,
		MPAARating:  "PG-13",
		ReleaseDate: rd,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, harryPotter)
	movies = append(movies, terminator)
	out, err := json.Marshal(movies)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}
