package menu

import (
	"strings"
	"testing"
	"zoo/models"
	"zoo/services"
)

func newTestZoo() *services.Zoo {
	return services.NewZoo(&mockClinic{})
}

type mockClinic struct{}

func (m *mockClinic) CheckHealth() bool { return true }

func TestAddTiger(t *testing.T) {
	z := newTestZoo()
	input := "1\nТигр\n5\n50\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addAnimalMenu()

	if len(z.Animals) != 1 {
		t.Fatalf("ожидалось 1 животное, получили %d", len(z.Animals))
	}

	a := z.Animals[0].(*models.Tiger)
	if a.GetName() != "Тигр" || a.GetFood() != 5 || a.Power != 50 {
		t.Fatalf("данные тигра не совпадают")
	}
}

func TestAddMonkey(t *testing.T) {
	z := newTestZoo()
	input := "3\nОбезьяна\n6\n7\n80\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addAnimalMenu()

	if len(z.Animals) != 1 {
		t.Fatalf("ожидалось 1 животное, получили %d", len(z.Animals))
	}

	a := z.Animals[0].(*models.Monkey)
	if a.GetName() != "Обезьяна" || a.GetFood() != 6 || a.Kindness != 7 || a.Intelligence != 80 {
		t.Fatalf("данные обезьяны не совпадают")
	}
}

func TestAddTable(t *testing.T) {
	z := newTestZoo()
	input := "1\nСтол\nДерево\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addThingMenu()

	if len(z.Things) != 1 {
		t.Fatalf("ожидалось 1 вещь, получили %d", len(z.Things))
	}

	tbl := z.Things[0].(models.Table)
	if tbl.Name != "Стол" || tbl.Material != "Дерево" {
		t.Fatalf("данные стола не совпадают")
	}
}

func TestAddComputer(t *testing.T) {
	z := newTestZoo()
	input := "2\nКомпьютер\nIntel\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addThingMenu()

	if len(z.Things) != 1 {
		t.Fatalf("ожидалось 1 вещь, получили %d", len(z.Things))
	}

	comp := z.Things[0].(models.Computer)
	if comp.Name != "Компьютер" || comp.CPU != "Intel" {
		t.Fatalf("данные компьютера не совпадают")
	}
}
