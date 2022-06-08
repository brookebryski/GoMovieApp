package service

import (
	"errors"
	"GoMovieApp/repository"
	"GoMovieApp/model"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrTitleIsNotEmpty = errors.New("Movie title cannot be empty")
	ErrMovieNotFound   = errors.New("the movie cannot be found")
)

type DefaultMovieService struct {
	movieRepo repository.IMovieRepository
}

func NewDefaultMovieService(mRepo repository.IMovieRepository) *DefaultMovieService {
	return &DefaultMovieService{
		movieRepo: mRepo,
	}
}


func (d *DefaultMovieService) GetMovies() ([]model.Movie, error) {
	return d.movieRepo.GetMovies()
}

func (d *DefaultMovieService) GetMovie(id int) (model.Movie, error) {
	if id <= 0 {
		return model.Movie{}, ErrIDIsNotValid
	}
	movie, err := d.movieRepo.GetMovie(id)

	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return model.Movie{}, ErrMovieNotFound
		}
	}
	return movie, nil
}
