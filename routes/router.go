package routes

import (
    "github.com/gorilla/mux"
    "github.com/SujinanFms/sujinan-backend/handlers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/api/contact", handlers.HandleContact).Methods("POST")
    r.HandleFunc("/api/recommendations", handlers.HandleRecommendations).Methods("GET", "POST")

    return r
}
