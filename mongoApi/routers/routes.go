package routes

import (
	"github.com/Shrut26/mongoApi/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controllers.DeleteAllMovie).Methods("DELETE")

	return r
}
