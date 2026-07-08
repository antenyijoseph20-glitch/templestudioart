package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}
}
func TestHomeHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d",
			http.StatusMethodNotAllowed,
			rr.Code)
	}
}
func TestHomeHandlerInvalidRoute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d",
			http.StatusNotFound,
			rr.Code)
	}
}
func TestDownloadHandler(t *testing.T) {

	form := url.Values{}
	form.Set("ascii", "HELLO")

	req := httptest.NewRequest(
		http.MethodPost,
		"/download",
		strings.NewReader(form.Encode()),
	)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	rr := httptest.NewRecorder()

	DownloadHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}

	if rr.Body.String() != "HELLO" {
		t.Errorf("Unexpected body: %q", rr.Body.String())
	}
}
func TestDownloadHandlerMethodNotAllowed(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/download",
		nil,
	)

	rr := httptest.NewRecorder()

	DownloadHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d got %d",
			http.StatusMethodNotAllowed,
			rr.Code)
	}
}
