package service

import (
	"encoding/json"
	"github.com/matthausen/bezero/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)
// TestSuggestionsCheckHandler -> check that the typeof response corresponds to what we expect
func TestSuggestionsCheckHandler(t *testing.T) {
	t.Run("Test SuggestionHandler endpoint", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/suggestions?q=London&latitude=51.507&longitude=0.127", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SuggestionHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var suggestions []model.Suggestion
		received := json.Unmarshal([]byte(rr.Body.String()), &suggestions)
		expected := `[{"Name":"Brighton","Latitude": 50.10,"Longitude": 0.12,"Score":0.68},
						{"Name":"Newcastle","Latitude": 54.65,"Longitude": 1.34,"Score":0.34},
						{"Name":"Edinburgh","Latitude": 55.00,"Longitude": 1.45,"Score":0.96}]`
		expectedJson := json.Unmarshal([]byte(expected), &suggestions)

		if reflect.TypeOf(received) != reflect.TypeOf(expectedJson) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
