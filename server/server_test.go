package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHitCounter struct {

}

func (hc *MockHitCounter) IncrementAndGetCounter() int {
	return 1
}

func TestMainPageHandler(t *testing.T) {
	srv := New(http.NewServeMux(),&MockHitCounter{})
	srv.routes()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.handleHits())
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"site_hit_counter":1}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

}
