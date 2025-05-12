package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestRoutes(t *testing.T) {
	tests := []struct {
		name string
		route string
	}{
		{
			name: "Test / route",
			route: "/",
		},
		{
			name: "Test /ping route",
			route: "/ping",
		},
		{
			name: "Test /cached route",
			route: "/cached",
		},
		{
			name: "Test /headers route",
			route: "/headers",
		},
		{
			name: "Test /env route",
			route: "/env",
		},
		{
			name: "Test /status route",
			route: "/status",
		},
	}

	router := setupRouter()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Logf("Running %s", test.name)

			req, _ := http.NewRequest("GET", test.route, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

		})
	}
}