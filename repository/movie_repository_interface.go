package repository

import (
	"GoMovieApp/model"
)

type IMovieRepository interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
}