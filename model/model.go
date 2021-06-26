package model

type (
	City struct {
		Name        string
		Latitude    float32
		Longitude   float32
		CountryCode string
	}

	Suggestion struct {
		Name      string
		Latitude  float32
		Longitude float32
		Score     float64
	}
)
