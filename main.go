package main 

import (
	"GoMovieApp/handler"
	"GoMovieApp/repository"
	"GoMovieApp/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

)

func main() {
	movieInMemoryRepository := repository.NewInMemoryMovieRepository()
	movieService := service.NewDefaultMovieService(movieInMemoryRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	router := httprouter.New()

	router.GET("/movies", movieHandler.GetMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)

	router.POST("/movies", movieHandler.CreateMovie)

	router.DELETE("/movies", movieHandler.DeleteAllMovies)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)

	router.PATCH("/movies/:id", movieHandler.UpdateMovie)

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)

}

// Handler: Layer that gets http request and returns http response to the client.
// Service: Layer that our business logic is in.
// Repository: Layer that provides all necessary data from external (DBs) or internal (in-memory) data source.