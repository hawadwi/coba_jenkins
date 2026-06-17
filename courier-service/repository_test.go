package main

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewDeliveryRepository(db)

	d := &Delivery{
		Resi:         "RESI001",
		CourierID:    1,
		AssignedZone: "A",
		Status:       "in_delivery",
	}

	mock.ExpectExec("INSERT INTO deliveries").
		WithArgs(
			d.Resi,
			d.CourierID,
			d.AssignedZone,
			d.Status,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(d)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByResi(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDeliveryRepository(db)

	now := time.Now()

	rows := sqlmock.NewRows(
		[]string{
			"resi",
			"courier_id",
			"assigned_zone",
			"status",
			"created_at",
		},
	).AddRow(
		"RESI001",
		1,
		"A",
		"in_delivery",
		now,
	)

	mock.ExpectQuery("SELECT").
		WithArgs("RESI001").
		WillReturnRows(rows)

	d, err := repo.GetByResi("RESI001")

	if err != nil {
		t.Fatal(err)
	}

	if d.Resi != "RESI001" {
		t.Fatal("wrong resi")
	}
}

func TestGetAll(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDeliveryRepository(db)

	now := time.Now()

	rows := sqlmock.NewRows(
		[]string{
			"resi",
			"courier_id",
			"assigned_zone",
			"status",
			"created_at",
		},
	).
		AddRow("R1", 1, "A", "in_delivery", now).
		AddRow("R2", 2, "B", "delivered", now)

	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)

	data, err := repo.GetAll()

	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 got %d", len(data))
	}
}