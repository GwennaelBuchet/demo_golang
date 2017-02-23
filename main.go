package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"zenika.com/demos/calc/calc"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTION", "PUT"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/calc/roman/sum/{a}/{b}", calc.GetRomanSum)
	router.HandleFunc("/api/v1/calc/arabic/sum/{a}/{b}", calc.GetArabicSum)

	handler := c.Handler(router)

	http.ListenAndServe(":8090", handler)
}
