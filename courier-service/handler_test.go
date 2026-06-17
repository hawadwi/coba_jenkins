package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockCourierService struct{}

func (m *mockCourierService) StartDelivery(d *Delivery) error {
	return nil
}

func (m *mockCourierService) CompleteDelivery(d *Delivery) error {
	return nil
}

func (m *mockCourierService) GetCourierDeliveries(
	deliveries []Delivery,
	courierID int,
) []Delivery {
	return deliveries
}

func (m *mockCourierService) ValidateDelivery(
	delivery *Delivery,
) error {
	return nil
}

func TestStartDelivery_InvalidBody(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/delivery",
		strings.NewReader("{"),
	)

	rec := httptest.NewRecorder()

	handler.StartDelivery(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestStartDelivery_MissingFields(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	body := `{
		"resi":"",
		"courier_id":0,
		"assigned_zone":""
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/delivery",
		strings.NewReader(body),
	)

	rec := httptest.NewRecorder()

	handler.StartDelivery(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestGetCourierDeliveries_EmptyID(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/courier/deliveries",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetCourierDeliveries(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestGetCourierDeliveries_InvalidID(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/courier/deliveries?courier_id=abc",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetCourierDeliveries(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestCompleteDelivery_InvalidBody(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/delivery/complete",
		strings.NewReader("{"),
	)

	rec := httptest.NewRecorder()

	handler.CompleteDelivery(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestHealth(t *testing.T) {

	handler := NewCourierHandler(
		&mockCourierService{},
		nil,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}
