package main

// import (
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

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
