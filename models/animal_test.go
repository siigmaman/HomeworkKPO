package models

import "testing"

func TestAnimal_GetSetFood(t *testing.T) {
	a := Animal{Name: "Leo", Food: 5, Number: 101}
	if a.GetFood() != 5 {
		t.Errorf("Expected Food=5, got %d", a.GetFood())
	}

	a.SetFood(10)
	if a.GetFood() != 10 {
		t.Errorf("Expected Food=10, got %d", a.GetFood())
	}

	if a.GetName() != "Leo" {
		t.Errorf("Expected Name=Leo, got %s", a.GetName())
	}

	if a.GetNumber() != 101 {
		t.Errorf("Expected Number=101, got %d", a.GetNumber())
	}
}

func TestHerbo_IsFriendly(t *testing.T) {
	h1 := Herbo{Animal: Animal{"Bunny", 2, 102}, Kindness: 4}
	h2 := Herbo{Animal: Animal{"Monkey", 3, 103}, Kindness: 6}
	if h1.IsFriendly() {
		t.Errorf("Expected h1 not friendly")
	}

	if !h2.IsFriendly() {
		t.Errorf("Expected h2 friendly")
	}
}

func TestTigerCreation(t *testing.T) {
	tiger := NewTiger("Tiger", 10, 80, 201)
	if tiger.Power != 80 {
		t.Errorf("Expected Power=80, got %d", tiger.Power)
	}

	if tiger.GetFood() != 10 {
		t.Errorf("Expected Food=10, got %d", tiger.GetFood())
	}
}

func TestWolfCreation(t *testing.T) {
	wolf := NewWolf("Wolf", 5, 70, 202)

	if wolf.Agility != 70 {
		t.Errorf("Expected Agility=70, got %d", wolf.Agility)
	}
}

func TestMonkeyCreation(t *testing.T) {
	monkey := NewMonkey("Momo", 4, 7, 90, 203)
	if monkey.Kindness != 7 || monkey.Intelligence != 90 {
		t.Errorf("Unexpected monkey fields")
	}

	if !monkey.IsFriendly() {
		t.Errorf("Expected friendly monkey")
	}
}

func TestRabbitCreation(t *testing.T) {
	rabbit := NewRabbit("Rabi", 3, 6, 204)

	if !rabbit.IsFriendly() {
		t.Errorf("Expected friendly rabbit")
	}
}

func TestThings(t *testing.T) {
	table := NewTable("Desk", 301, "Wood")
	if table.GetNumber() != 301 {
		t.Errorf("Expected 301, got %d", table.GetNumber())
	}

	computer := NewComputer("PC", 302, "Intel")
	if computer.GetNumber() != 302 {
		t.Errorf("Expected 302, got %d", computer.GetNumber())
	}
}
