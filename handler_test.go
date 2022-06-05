package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandler(t *testing.T) {

	router := setupRouter(false, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	output := ExternalAdapterOutput{}
	json.NewDecoder(w.Body).Decode(&output)

	assert.NotNil(t, output.Data)
	assert.True(t, output.Data.Price > 0)

}
