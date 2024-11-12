package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "The Lord of the Rings")
}

func TestGetMovie(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		movieID    string
		wantStatus int
		wantBody   string
	}{
		{"Valid ID 1", "1", http.StatusOK, "The Lord of the Rings"},
		{"Valid ID 2", "2", http.StatusOK, "Inception"},
		{"Invalid ID", "999", http.StatusNotFound, "Movie not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/movies/"+tt.movieID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestCreateMovie(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		input      string
		wantStatus int
		wantBody   string
	}{
		{
			name: "Valid Movie",
			input: `{
            "isbn": "9999",
            "title": "New Movie",
            "director": {
                "firstname": "New",
                "lastname": "Director"
            }
        }`,
			wantStatus: http.StatusCreated,
			wantBody:   "New Movie",
		},
		{
			name: "Invalid Movie",
			input: `{
            "isbn": 9999,
            "title": "",
            "director": {
                "firstname": "",
                "lastname": ""
            }
        }`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/movies", bytes.NewBufferString(tt.input))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		movieID    string
		input      string
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Valid Update",
			movieID: "3",
			input: `{
                "isbn": "5678",
                "title": "The Matrix Reloaded",
                "director": {
                    "firstname": "Lana",
                    "lastname": "Wachowski"
                }
            }`,
			wantStatus: http.StatusOK,
			wantBody:   "The Matrix Reloaded",
		},
		{
			name:    "Invalid Update",
			movieID: "3",
			input: `{
                "isbn": 999,
                "title": "",
                "director": {
                    "firstname": "",
                    "lastname": ""
                }
            }`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "error",
		},
		{
			name:    "Non-existent Movie",
			movieID: "999",
			input: `{
                "isbn": "5678",
                "title": "Non-existent Movie",
                "director": {
                    "firstname": "Non",
                    "lastname": "Existent"
                }
            }`,
			wantStatus: http.StatusNotFound,
			wantBody:   "Movie not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("PUT", "/movies/"+tt.movieID, bytes.NewBufferString(tt.input))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.wantBody)
		})
	}
}

func TestDeleteMovie(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		movieID    string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Valid Delete",
			movieID:    "1",
			wantStatus: http.StatusNoContent,
			wantBody:   "",
		},
		{
			name:       "Non-existent Movie",
			movieID:    "999",
			wantStatus: http.StatusNotFound,
			wantBody:   "Movie not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/movies/"+tt.movieID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			if tt.wantBody != "" {
				assert.Contains(t, w.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestServeIndex(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Movie CRUD Operations")
}
