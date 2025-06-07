package handlers

import (
	"cutURL/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateURLHandler(t *testing.T) {
	type want struct {
		statusCode int
		body       string
	}

	tests := []struct {
		name  string
		input string
		want  want
	}{
		{
			name:  "success",
			input: "https://google.com",
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name:  "invalid URL",
			input: "abc",
			want: want{
				statusCode: http.StatusBadRequest,
				body:       `{"error":"Invalid URL"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			router := gin.Default()
			storage := storage.NewStorage()
			baseURLTest := "http://localhost:8080"
			router.POST("/create", CreateURLHandler(storage, baseURLTest))

			req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(tt.input))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != tt.want.statusCode {
				t.Errorf("StatusCode = %v, want %v", resp.Code, tt.want.statusCode)
			}

			if tt.want.body != "" && strings.TrimSpace(resp.Body.String()) != tt.want.body {
				t.Errorf("Response body = %v, want %v", resp.Body.String(), tt.want.body)
			}
		})
	}
}

func TestGetURLHandler(t *testing.T) {
	type want struct {
		statusCode int
		location   string
		body       string
	}

	tests := []struct {
		name      string
		id        string
		setupData map[string]string
		want      want
	}{
		{
			name: "success redirect",
			id:   "abc123",
			setupData: map[string]string{
				"abc123": "https://google.com",
			},
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				location:   "https://google.com",
			},
		},
		{
			name:      "not found",
			id:        "notexist",
			setupData: map[string]string{},
			want: want{
				statusCode: http.StatusNotFound,
				body:       `{"error":"URL not found"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			router := gin.Default()
			storage := storage.NewStorage()

			for k, v := range tt.setupData {
				storage.SetData(v, k)
			}

			router.GET("/:id", GetURLHandler(storage))

			req := httptest.NewRequest(http.MethodGet, "/"+tt.id, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != tt.want.statusCode {
				t.Errorf("StatusCode = %v, want %v", resp.Code, tt.want.statusCode)
			}

			if tt.want.location != "" {
				location := resp.Header().Get("Location")
				if location != tt.want.location {
					t.Errorf("Location = %v, want %v", location, tt.want.location)
				}
			}

			if tt.want.body != "" {
				if strings.TrimSpace(resp.Body.String()) != tt.want.body {
					t.Errorf("Body = %v, want %v", resp.Body.String(), tt.want.body)
				}
			}
		})
	}
}
