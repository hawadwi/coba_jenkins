package main

import "testing"

func TestStartDelivery_Success(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		Resi:         "RESI001",
		CourierID:    1,
		Status:       "pending",
		AssignedZone: "A",
	}

	err := service.StartDelivery(d)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.Status != "in_delivery" {
		t.Fatalf("expected in_delivery got %s", d.Status)
	}
}

func TestStartDelivery_NilDelivery(t *testing.T) {

	service := &CourierService{}

	err := service.StartDelivery(nil)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestStartDelivery_EmptyResi(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		CourierID: 1,
		Status:    "pending",
	}

	err := service.StartDelivery(d)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestStartDelivery_InvalidCourier(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		Resi:      "RESI001",
		CourierID: 0,
		Status:    "pending",
	}

	err := service.StartDelivery(d)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCompleteDelivery_Success(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		Resi:   "RESI001",
		Status: "in_delivery",
	}

	err := service.CompleteDelivery(d)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.Status != "delivered" {
		t.Fatalf("expected delivered got %s", d.Status)
	}

	if d.DeliveredAt == nil {
		t.Fatal("expected delivered_at")
	}
}

func TestCompleteDelivery_InvalidStatus(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		Status: "pending",
	}

	err := service.CompleteDelivery(d)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetCourierDeliveries(t *testing.T) {

	service := &CourierService{}

	data := []Delivery{
		{Resi: "A", CourierID: 1},
		{Resi: "B", CourierID: 2},
		{Resi: "C", CourierID: 1},
	}

	result := service.GetCourierDeliveries(data, 1)

	if len(result) != 2 {
		t.Fatalf("expected 2 got %d", len(result))
	}
}

func TestValidateDelivery_Success(t *testing.T) {

	service := &CourierService{}

	d := &Delivery{
		Resi:           "RESI001",
		CourierID:      1,
		AlamatPenerima: "Bandung",
	}

	err := service.ValidateDelivery(d)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateDelivery_Nil(t *testing.T) {

	service := &CourierService{}

	err := service.ValidateDelivery(nil)

	if err == nil {
		t.Fatal("expected error")
	}
}
