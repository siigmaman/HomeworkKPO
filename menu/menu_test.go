package menu

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
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

func TestAddWolf(t *testing.T) {
	z := newTestZoo()
	input := "2\nВолк\n4\n75\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addAnimalMenu()

	if len(z.Animals) != 1 {
		t.Fatalf("ожидалось 1 животное, получили %d", len(z.Animals))
	}

	a := z.Animals[0].(*models.Wolf)
	if a.GetName() != "Волк" || a.GetFood() != 4 || a.Agility != 75 {
		t.Fatalf("данные волка не совпадают")
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

func TestAddRabbit(t *testing.T) {
	z := newTestZoo()
	input := "4\nКролик\n2\n8\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	m.addAnimalMenu()

	if len(z.Animals) != 1 {
		t.Fatalf("ожидалось 1 животное, получили %d", len(z.Animals))
	}

	a := z.Animals[0].(*models.Rabbit)
	if a.GetName() != "Кролик" || a.GetFood() != 2 || a.Kindness != 8 {
		t.Fatalf("данные кролика не совпадают")
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

func TestShowContactZoo_WithFriendlyAnimals(t *testing.T) {
	z := newTestZoo()
	monkey := models.NewMonkey("Дружелюбная обезьяна", 5, 8, 70, 1001)
	rabbit := models.NewRabbit("Дружелюбный кролик", 2, 9, 1002)
	z.AddAnimal(&monkey)
	z.AddAnimal(&rabbit)

	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showContactZoo()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Дружелюбная обезьяна") {
		t.Error("ожидалось увидеть дружелюбную обезьяну в выводе")
	}
	if !strings.Contains(output, "Дружелюбный кролик") {
		t.Error("ожидалось увидеть дружелюбного кролика в выводе")
	}
}

func TestShowContactZoo_NoFriendlyAnimals(t *testing.T) {
	z := newTestZoo()
	tiger := models.NewTiger("Тигр", 10, 80, 1001)
	z.AddAnimal(&tiger)

	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showContactZoo()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Нет животных") {
		t.Error("ожидалось сообщение об отсутствии подходящих животных")
	}
}

func TestShowAllInventory_Empty(t *testing.T) {
	z := newTestZoo()
	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showAllInventory()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Нет животных") {
		t.Error("ожидалось сообщение об отсутствии животных")
	}
	if !strings.Contains(output, "Нет вещей") {
		t.Error("ожидалось сообщение об отсутствии вещей")
	}
}

func TestShowAllInventory_WithItems(t *testing.T) {
	z := newTestZoo()
	monkey := models.NewMonkey("Обезьяна", 5, 8, 70, 1001)
	table := models.NewTable("Стол", 2001, "Дерево")
	z.AddAnimal(&monkey)
	z.AddThing(table)

	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showAllInventory()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Обезьяна") {
		t.Error("ожидалось увидеть обезьяну в выводе")
	}
	if !strings.Contains(output, "Стол") {
		t.Error("ожидалось увидеть стол в выводе")
	}
	if !strings.Contains(output, "Итого на балансе: 2 предметов") {
		t.Error("ожидалась корректная статистика предметов")
	}
}

func TestFeedAnimal(t *testing.T) {
	z := newTestZoo()
	tiger := models.NewTiger("Тигр", 10, 80, 1001)
	z.AddAnimal(&tiger)

	input := "1\n5\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.feedAnimal()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if tiger.GetFood() != 15 {
		t.Errorf("ожидалось 15 кг еды, получили %d", tiger.GetFood())
	}
	if !strings.Contains(output, "15 кг/день") {
		t.Error("ожидалось сообщение об обновленном количестве еды")
	}
}

func TestFeedAnimal_NoAnimals(t *testing.T) {
	z := newTestZoo()
	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.feedAnimal()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Нет животных") {
		t.Error("ожидалось сообщение об отсутствии животных для кормления")
	}
}

func TestCheckAnimalHealth(t *testing.T) {
	z := newTestZoo()
	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.checkAnimalHealth()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Животное здорово!") {
		t.Error("ожидалось сообщение о здоровом животном")
	}
}

func TestShowTemperature(t *testing.T) {
	z := newTestZoo()
	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showTemperature()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Температура в зоопарке:") {
		t.Error("ожидалось сообщение о температуре")
	}
}

func TestShowAllAnimalsDetailed_Empty(t *testing.T) {
	z := newTestZoo()
	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showAllAnimalsDetailed()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Животных пока нет") {
		t.Error("ожидалось сообщение об отсутствии животных")
	}
}

func TestShowAllAnimalsDetailed_WithAnimals(t *testing.T) {
	z := newTestZoo()
	tiger := models.NewTiger("Тигр", 10, 80, 1001)
	monkey := models.NewMonkey("Обезьяна", 5, 8, 70, 1002)
	z.AddAnimal(&tiger)
	z.AddAnimal(&monkey)

	m := NewMenuWithReader(z, strings.NewReader(""))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.showAllAnimalsDetailed()

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Тигр") || !strings.Contains(output, "Обезьяна") {
		t.Error("ожидалось увидеть всех животных в выводе")
	}
	if !strings.Contains(output, "Сила: 80") {
		t.Error("ожидалось увидеть параметры тигра")
	}
	if !strings.Contains(output, "Доброта: 8") || !strings.Contains(output, "Интеллект: 70") {
		t.Error("ожидалось увидеть параметры обезьяны")
	}
}

func TestReadInt_ValidInput(t *testing.T) {
	z := newTestZoo()
	input := "5\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	result := m.readInt("Введите число: ", 1, 10)

	if result != 5 {
		t.Errorf("ожидалось 5, получили %d", result)
	}
}

func TestReadInt_InvalidThenValidInput(t *testing.T) {
	z := newTestZoo()
	input := "abc\n15\n5\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	result := m.readInt("Введите число: ", 1, 10)

	if result != 5 {
		t.Errorf("ожидалось 5, получили %d", result)
	}
}

func TestMenuRun_Exit(t *testing.T) {
	z := newTestZoo()
	input := "0\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	done := make(chan bool)
	go func() {
		m.Run()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("меню не завершилось при выборе выхода")
	}
}

func TestMenuRun_InvalidChoice(t *testing.T) {
	z := newTestZoo()
	input := "99\n0\n"
	m := NewMenuWithReader(z, strings.NewReader(input))

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan bool)
	go func() {
		m.Run()
		done <- true
	}()

	time.Sleep(100 * time.Millisecond)

	m.scanner = bufio.NewScanner(strings.NewReader("0\n"))

	<-done

	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Неверный выбор") {
		t.Error("ожидалось сообщение о неверном выборе")
	}
}
