package service

import (
	"fmt"
	"testing"
)

// London to Manchester is ~262km. We'll check with a +- 10km precision
// Rome to Moscow is ~ 2376.612. We'll check with a +- 10km precision
func TestHaversineDistance(t *testing.T) {
	t.Run("Should calculate the right distance in km between 2 coordinates lat long", func(t *testing.T) {
		latManchester := 53.48
		longManchester := 2.24

		latLondon := 51.50
		longLondon := 0.12

		distance :=  HaversineDistance(latManchester, longManchester, latLondon, longLondon)
		fmt.Printf("Distance calculated : %v", distance)
		expected := 262.4

		if distance < expected - 10 && distance > expected + 10 {
			t.Errorf("Wrong distance in km calculated: got %v want %v",
				distance, expected)
		}
	})

	t.Run("Should calculate the right distance in km between 2 coordinates lat long", func(t *testing.T) {
		latRome := 41.90
		longRome := 12.49

		latMoscow := 55.75
		longMoscow := 37.61

		distance :=  HaversineDistance(latRome, longRome, latMoscow, longMoscow)
		fmt.Printf("Distance calculated : %v", distance)
		expected := 2376.61

		if distance < expected - 10 && distance > expected + 10 {
			t.Errorf("Wrong distance in km calculated: got %v want %v",
				distance, expected)
		}
	})
}


// TestReadCsv -> check that we can open a csv file
func TestCSVReader(t *testing.T){
	t.Run("Should open a csv file and read from it", func(t *testing.T) {
		row := ReadCsvFile("../cities_cleaned.csv")

		rows := len(row)
		expectedRows := 769
		if rows != expectedRows {
			t.Errorf("Row Length is %d, Expected %d", rows, expectedRows)
		}
	})
}
