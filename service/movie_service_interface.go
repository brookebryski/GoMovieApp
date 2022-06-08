package service

import "GoMovieApp/model"

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
}