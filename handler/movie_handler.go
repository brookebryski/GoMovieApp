package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"GoMovieApp/service"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type movieHandler struct {
	service service.IMovieService
}

func NewMovieHandler(ms service.IMovieService) *movieHandler {
	return &movieHandler{service: ms}
}

// curl localhost:8080/movies | jq
func (mh *movieHandler) GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := mh.service.GetMovies()
	if err != nil {
		http.Error(w, "Unable to get all movies", http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

// curl "localhost:8080/movies/1" | jq
func (mh *movieHandler) GetMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	movie, err := mh.service.GetMovie(id)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrMovieNotFound) { // Test yaz
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}
