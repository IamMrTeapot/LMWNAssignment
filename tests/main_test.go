package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IamMrTeapot/LWMNAssignment/server"
	"github.com/stretchr/testify/assert"
)

func TestCovidSummary(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
