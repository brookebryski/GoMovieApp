package handler

import (
	"encoding/json"
	"net/http"
	"GoMovieApp/service"
	"github.com/julienschmidt/httprouter"
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

