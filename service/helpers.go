package service

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/matthausen/bezero/model"
)

// ReadCsvFile -> given a file path, return the records
func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// FetchAllCities -> returns a comlpete list of GB cities we can later sort from
func FetchAllCities() *[]model.City {
	records := ReadCsvFile("./cities_cleaned.csv")
	var allCities []model.City
	for _, value := range records {
		lat, err := strconv.ParseFloat(value[1], 32)
		if err != nil {
			log.Printf("Could not convert string latitude to float64: %v\n", err)
		}

		long, err := strconv.ParseFloat(value[2], 32)
		if err != nil {
			log.Printf("Could not convert string latitude to float64: %v\n", err)
		}
		city := model.City{Name: value[0], Latitude: float32(lat), Longitude: float32(long), CountryCode: value[3]}
		allCities = append(allCities, city)
	}
	return &allCities
}

// HaversineDistance -> returns the distance in km between 2 coordinates
func HaversineDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	R := 6371 // Radius of the earth in km
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(float64(dLat/2))*math.Sin(float64(dLat/2)) +
		math.Cos(float64(deg2rad(lat1)))*math.Cos(float64(deg2rad(lat2)))*
			math.Sin(float64(dLon/2))*math.Sin(float64(dLon/2))

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := float64(R) * c
	return d
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
