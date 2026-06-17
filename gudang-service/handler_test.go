package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStartSort_InvalidBody(t *testing.T) {

	handler := &SortingHandler{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/sort",
		strings.NewReader("{"),
	)

	rec := httptest.NewRecorder()

	handler.StartSort(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestCompleteSort_InvalidBody(t *testing.T) {

	handler := &SortingHandler{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/sort/complete",
		strings.NewReader("{"),
	)

	rec := httptest.NewRecorder()

	handler.CompleteSort(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
