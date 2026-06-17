package main

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreatePackage(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	pkg := &Package{
		Resi:          "RESI001",
		NamaBarang:    "Laptop",
		Berat:         2,
		WarehouseZone: "A",
		Status:        "sorting",
	}

	mock.ExpectExec("INSERT INTO packages").
		WithArgs(
			pkg.Resi,
			pkg.NamaBarang,
			pkg.Berat,
			pkg.WarehouseZone,
			pkg.Status,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(pkg)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetByResi(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	now := time.Now()

	rows := sqlmock.NewRows(
		[]string{
			"resi",
			"nama_barang",
			"berat",
			"warehouse_zone",
			"status",
			"sorted_at",
		},
	).AddRow(
		"RESI001",
		"Laptop",
		2,
		"A",
		"sorting",
		now,
	)

	mock.ExpectQuery("SELECT").
		WithArgs("RESI001").
		WillReturnRows(rows)

	pkg, err := repo.GetByResi("RESI001")

	if err != nil {
		t.Fatal(err)
	}

	if pkg.Resi != "RESI001" {
		t.Fatal("wrong resi")
	}
}

func TestCompleteSort(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	mock.ExpectExec("UPDATE packages").
		WithArgs("RESI001").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.CompleteSort("RESI001")

	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveOutbox(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	mock.ExpectExec("INSERT INTO outbox_events").
		WithArgs(
			"PACKAGE_READY",
			`{"resi":"RESI001"}`,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.SaveOutbox(
		"PACKAGE_READY",
		`{"resi":"RESI001"}`,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateStatus(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	mock.ExpectExec("UPDATE packages SET status").
		WithArgs(
			"ready",
			"RESI001",
		).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateStatus(
		"RESI001",
		"ready",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestMarkAsSent(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	mock.ExpectExec("UPDATE outbox_events").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.MarkAsSent(1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPendingEvents(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	rows := sqlmock.NewRows(
		[]string{
			"id",
			"event_type",
			"payload",
		},
	).AddRow(
		1,
		"PACKAGE_READY",
		`{"resi":"RESI001"}`,
	)

	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)

	data, err := repo.GetPendingEvents()

	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 got %d", len(data))
	}
}

func TestGetAllPackages(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPackageRepository(db)

	now := time.Now()

	rows := sqlmock.NewRows(
		[]string{
			"resi",
			"nama_barang",
			"berat",
			"warehouse_zone",
			"status",
			"sorted_at",
		},
	).AddRow(
		"R1",
		"Laptop",
		2,
		"A",
		"sorting",
		now,
	)

	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)

	data, err := repo.GetAll()

	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 got %d", len(data))
	}
}