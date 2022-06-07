package repository

import (
	"errors"
	"GoMovieApp/model"
)

var (
	ErrMovieNotFound = errors.New("FromRepository - movie not found")
)

type inmemoryMovieRepository struct {
	Movies []model.Movie
}

func NewInMemoryMovieRepository() *inmemoryMovieRepository {
	var movies = []model.Movie{
		{ID: 1, Title: "The Shawshank Redemption", ReleaseYear: 1994, Score: 9.3},
		{ID: 2, Title: "The Godfather", ReleaseYear: 1972, Score: 9.2},
		{ID: 3, Title: "The Dark Knight", ReleaseYear: 2008, Score: 9.0},
	}
	return &inmemoryMovieRepository{
		Movies: movies,
	}
}

func (i *inmemoryMovieRepository) GetMovies() ([]model.Movie, error) {
	return i.Movies, nil
}