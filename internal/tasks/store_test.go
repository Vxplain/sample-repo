package tasks

import (
	"errors"
	"testing"
)

func TestStoreAddAndList(t *testing.T) {
	store := NewStore()

	first, err := store.Add("write tests")
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}
	second, err := store.Add("review pull request")
	if err != nil {
		t.Fatalf("Add() error = %v", err)
	}

	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("IDs = %d, %d; want 1, 2", first.ID, second.ID)
	}

	got := store.List()
	if len(got) != 2 {
		t.Fatalf("List() length = %d; want 2", len(got))
	}

	got[0].Title = "changed copy"
	if store.List()[0].Title != "write tests" {
		t.Fatal("List() exposed mutable store state")
	}
}

func TestStoreRejectsEmptyTitle(t *testing.T) {
	store := NewStore()

	_, err := store.Add("  ")
	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("Add() error = %v; want %v", err, ErrEmptyTitle)
	}
}
