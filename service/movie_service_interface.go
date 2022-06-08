package service

import "GoMovieApp/model"

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	UpdateMovie(id int, movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovies() error
}