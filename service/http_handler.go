package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/matthausen/bezero/model"
)

// SuggestionHandler -> The endpoint which returns suggestions
func SuggestionHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("q")
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	fmt.Printf("Search for: %s, %s, %s\n", city, latitude, longitude)

	lat, err := strconv.ParseFloat(latitude, 32)
	if err != nil {
		log.Printf("Could not convert string latitude to float64: %v\n", err)
	}

	long, err := strconv.ParseFloat(longitude, 32)
	if err != nil {
		log.Printf("Could not convert string latitude to float64: %v\n", err)
	}

	allcities := FetchAllCities()

	var partialSuggestions []model.Suggestion
	var suggestions []model.Suggestion

	var distances float64
	for _, c := range *allcities {
		var distance float64

		if strings.Contains(c.Name, city) {
			distance = HaversineDistance(float64(c.Latitude), float64(c.Longitude), lat, long)
			distances += distance
			partialSuggestions = append(partialSuggestions, model.Suggestion{
				Name:      c.Name,
				Latitude:  c.Latitude,
				Longitude: c.Longitude,
			})
		}
	}

	for _, s := range partialSuggestions {
		distance := HaversineDistance(float64(s.Latitude), float64(s.Longitude), lat, long)
		x := (distance * 100) / distances
		suggestions = append(suggestions, model.Suggestion{
			Name:      s.Name,
			Latitude:  s.Latitude,
			Longitude: s.Longitude,
			Score:     (100 - x) / 100,
		})
	}

	// fmt.Printf("Suggestions: %v\n", suggestions)
	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Score > suggestions[j].Score
	})

	json.NewEncoder(w).Encode(suggestions)
}
