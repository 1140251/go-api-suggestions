package service

import (
	"fmt"
	"net/http"
)

func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("q")
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	fmt.Printf("Search for: %s, %s, %s\n", city, latitude, longitude)
}
