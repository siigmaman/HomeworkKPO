package services

import (
	"testing"
	"zoo/models"
)

type mockClinic struct{}

func (m *mockClinic) CheckHealth() bool { return true }

func TestZoo_AddAnimalAndThing(t *testing.T) {
	clinic := &mockClinic{}
	z := NewZoo(clinic)

	tiger := models.NewTiger("Tiger", 5, 80, 101)
	z.AddAnimal(&tiger)
	if len(z.Animals) != 1 {
		t.Errorf("Expected 1 animal, got %d", len(z.Animals))
	}

	table := models.NewTable("Desk", 201, "Wood")
	z.AddThing(table)
	if len(z.Things) != 1 {
		t.Errorf("Expected 1 thing, got %d", len(z.Things))
	}
}

func TestZoo_Report(t *testing.T) {
	clinic := &mockClinic{}
	z := NewZoo(clinic)
	tiger := models.NewTiger("Tiger", 5, 80, 101)
	z.AddAnimal(&tiger)
	z.Report()
}
