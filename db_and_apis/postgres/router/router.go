package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock", middleware.GetStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}
