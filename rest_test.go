package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := setupRest()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rest/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Rest Service is Up \\u0026 Running\"}", w.Body.String())

}

func TestPostToDoRoute(t *testing.T) {
	r := setupRest()
	w := httptest.NewRecorder()

	var readrForIo io.Reader
	readrForIo = strings.NewReader("{\"task\": \"one\", \"status\": \"sadad\"}")
	reqp, _ := http.NewRequest("POST", "/rest/task/", readrForIo)
	r.ServeHTTP(w, reqp)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"message\":\"Task Logged\"}", w.Body.String())
}

func TestGetToDoRoute(t *testing.T) {
	r := setupRest()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rest/task/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Body.String())

}

func TestGetSingleToDoRoute(t *testing.T) {
	r := setupRest()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rest/task/99999999999", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Equal(t, "{\"message\":\"Empty Result\"}", w.Body.String())
}
