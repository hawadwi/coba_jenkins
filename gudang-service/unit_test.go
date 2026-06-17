package main

import "testing"

func TestStartSorting_Success(t *testing.T) {
	svc := &SortingService{}

	pkg := &Package{
		Resi: "RESI001",
	}

	err := svc.StartSorting(pkg)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if pkg.Status != "sorting" {
		t.Fatalf("expected sorting, got %s", pkg.Status)
	}
}

func TestStartSorting_EmptyResi(t *testing.T) {
	svc := &SortingService{}

	pkg := &Package{
		Resi: "",
	}

	err := svc.StartSorting(pkg)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCompleteSorting_NilPackage(t *testing.T) {
	svc := &SortingService{}

	err := svc.CompleteSorting(nil)

	if err == nil {
		t.Fatal("expected error")
	}
}
