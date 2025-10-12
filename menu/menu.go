package menu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"zoo/models"
	"zoo/services"
)

type Menu struct {
	zoo              *services.Zoo
	scanner          *bufio.Scanner
	nextAnimalNumber int
	nextThingNumber  int
}

func NewMenu(zoo *services.Zoo) *Menu {
	return &Menu{
		zoo:              zoo,
		scanner:          bufio.NewScanner(os.Stdin),
		nextAnimalNumber: 1001,
		nextThingNumber:  2001,
	}
}

func NewMenuWithReader(zoo *services.Zoo, r io.Reader) *Menu {
	return &Menu{
		zoo:              zoo,
		scanner:          bufio.NewScanner(r),
		nextAnimalNumber: 1001,
		nextThingNumber:  2001,
	}
}

func (m *Menu) Run() {
	for {
		m.showMainMenu()
		m.scanner.Scan()
		choice := m.scanner.Text()

		switch choice {
		case "1":
			m.addAnimalMenu()

		case "2":
			m.addThingMenu()

		case "3":
			m.showReport()

		case "4":
			m.showContactZoo()

		case "5":
			m.showAllInventory()

		case "6":
			m.feedAnimal()

		case "7":
			m.checkAnimalHealth()

		case "8":
			m.showTemperature()

		case "9":
			m.showAllAnimalsDetailed()

		case "0":
			fmt.Println("Выход из программы...")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func (m *Menu) showMainMenu() {
	fmt.Println("\n===== Главное меню =====")
	fmt.Println("1. Добавить животное")
	fmt.Println("2. Добавить вещь в инвентарь")
	fmt.Println("3. Показать полный отчет")
	fmt.Println("4. Показать животных для контактного зоопарка")
	fmt.Println("5. Показать весь инвентарь")
	fmt.Println("6. Покормить животное")
	fmt.Println("7. Проверить здоровье животного")
	fmt.Println("8. Проверить температуру")
	fmt.Println("9. Показать всех животных с параметрами")
	fmt.Println("0. Выход")
	fmt.Print("Ваш выбор: ")
}

func (m *Menu) readInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		m.scanner.Scan()
		text := m.scanner.Text()
		value, err := strconv.Atoi(text)

		if err != nil || value < min || value > max {
			fmt.Printf("Введите корректное число от %d до %d.\n", min, max)
			continue
		}

		return value
	}
}

func (m *Menu) addAnimalMenu() {
	fmt.Println("\nДОБАВЛЕНИЕ НОВОГО ЖИВОТНОГО")
	fmt.Println("=============================")

	fmt.Println("Выберите тип животного:")
	fmt.Println("1. Тигр (хищник)")
	fmt.Println("2. Волк (хищник)")
	fmt.Println("3. Обезьяна (травоядное)")
	fmt.Println("4. Кролик (травоядное)")
	animalType := m.readInt("Ваш выбор: ", 1, 4)

	fmt.Print("Введите имя животного: ")
	m.scanner.Scan()
	name := m.scanner.Text()

	food := m.readInt("Введите потребление еды (кг/день): ", 1, 1000)

	var animal models.IAlive

	switch animalType {
	case 1:
		power := m.readInt("Введите силу тигра (1-100): ", 1, 100)
		t := models.NewTiger(name, food, power, m.nextAnimalNumber)
		animal = &t

	case 2:
		agility := m.readInt("Введите ловкость волка (1-100): ", 1, 100)
		t := models.NewWolf(name, food, agility, m.nextAnimalNumber)
		animal = &t

	case 3:
		kindness := m.readInt("Введите доброту обезьяны (1-10): ", 1, 10)
		intelligence := m.readInt("Введите интеллект обезьяны (1-100): ", 1, 100)
		t := models.NewMonkey(name, food, kindness, intelligence, m.nextAnimalNumber)
		animal = &t

	case 4:
		kindness := m.readInt("Введите доброту кролика (1-10): ", 1, 10)
		t := models.NewRabbit(name, food, kindness, m.nextAnimalNumber)
		animal = &t

	}

	m.zoo.AddAnimal(animal)
	m.nextAnimalNumber++
}

func (m *Menu) addThingMenu() {
	fmt.Println("\nДОБАВЛЕНИЕ ВЕЩИ В ИНВЕНТАРЬ")
	fmt.Println("=============================")

	fmt.Println("Выберите тип вещи:")
	fmt.Println("1. Стол")
	fmt.Println("2. Компьютер")
	thingType := m.readInt("Ваш выбор: ", 1, 2)

	fmt.Print("Введите название вещи: ")
	m.scanner.Scan()
	name := m.scanner.Text()

	var thing models.IInventory

	switch thingType {
	case 1:
		fmt.Print("Введите материал стола: ")
		m.scanner.Scan()
		material := m.scanner.Text()
		thing = models.NewTable(name, m.nextThingNumber, material)

	case 2:
		fmt.Print("Введите процессор компьютера: ")
		m.scanner.Scan()
		cpu := m.scanner.Text()
		thing = models.NewComputer(name, m.nextThingNumber, cpu)

	}

	m.zoo.AddThing(thing)
	m.nextThingNumber++
}

func (m *Menu) showReport() {
	fmt.Println("\nПОЛНЫЙ ОТЧЕТ ЗООПАРКА")
	fmt.Println("========================")
	m.zoo.Report()
}

func (m *Menu) showContactZoo() {
	fmt.Println("\nЖИВОТНЫЕ ДЛЯ КОНТАКТНОГО ЗООПАРКА")
	fmt.Println("==================================")

	animals := m.zoo.Animals
	hasFriendly := false

	for _, a := range animals {
		if f, ok := a.(interface{ IsFriendly() bool }); ok && f.IsFriendly() {
			inventoryItem := a.(models.IInventory)

			fmt.Printf(" %s (Инв. №%d) - %d кг/день\n",
				a.GetName(), inventoryItem.GetNumber(), a.GetFood())
			hasFriendly = true
		}
	}

	if !hasFriendly {
		fmt.Println("Нет животных, подходящих для контактного зоопарка")
		fmt.Println("(нужна доброта > 5 баллов)")
	}
}

func (m *Menu) showAllInventory() {
	fmt.Println("\nВЕСЬ ИНВЕНТАРЬ ЗООПАРКА")
	fmt.Println("=========================")

	fmt.Println("\nЖИВОТНЫЕ:")
	animals := m.zoo.Animals
	if len(animals) == 0 {
		fmt.Println("   Нет животных")
	} else {
		for _, a := range animals {
			inventoryItem := a.(models.IInventory)
			animalType := "Хищник"

			if f, ok := a.(interface{ IsFriendly() bool }); ok && f.IsFriendly() {
				animalType = "Травоядное"
			}

			fmt.Printf("   %s (№%d) - %s - %d кг/день\n",
				a.GetName(), inventoryItem.GetNumber(), animalType, a.GetFood())
		}
	}

	fmt.Println("\nВЕЩИ:")
	things := m.zoo.Things

	if len(things) == 0 {
		fmt.Println("   Нет вещей")
	} else {
		for _, t := range things {
			fmt.Printf("   %v (Инв. №%d)\n", t, t.GetNumber())
		}
	}

	totalItems := len(animals) + len(things)
	fmt.Printf("\nИтого на балансе: %d предметов\n", totalItems)
}

func (m *Menu) feedAnimal() {
	animals := m.zoo.Animals

	if len(animals) == 0 {
		fmt.Println("Нет животных для кормления.")
		return
	}

	fmt.Println("\nВыберите животное для кормления:")

	for i, a := range animals {
		fmt.Printf("%d. %s (ест %d кг/день)\n", i+1, a.GetName(), a.GetFood())
	}

	idx := m.readInt("Ваш выбор: ", 1, len(animals))
	animal := animals[idx-1]

	delta := m.readInt("Введите количество еды для добавления (кг): ", 1, 1000)
	animal.SetFood(animal.GetFood() + delta)
	fmt.Printf("%s теперь ест %d кг/день\n", animal.GetName(), animal.GetFood())
}

func (m *Menu) checkAnimalHealth() {
	healthy := m.zoo.Clinic.CheckHealth()

	if healthy {
		fmt.Println("Животное здорово!")
	} else {
		fmt.Println("Животное больно.")
	}
}

func (m *Menu) showTemperature() {
	temp := services.CheckTemp()
	fmt.Printf("Температура в зоопарке: %.1f°C\n", temp)
}

func (m *Menu) showAllAnimalsDetailed() {
	fmt.Println("\nВСЕ ЖИВОТНЫЕ ЗООПАРКА")
	fmt.Println("=======================")

	if len(m.zoo.Animals) == 0 {
		fmt.Println("Животных пока нет.")
		return
	}

	for _, a := range m.zoo.Animals {
		inventoryItem := a.(models.IInventory)
		animalType := "Хищник"

		if f, ok := a.(interface{ IsFriendly() bool }); ok && f.IsFriendly() {
			animalType = "Травоядное"
		}

		fmt.Printf("Имя: %s | №%d | Тип: %s | Потребление еды: %d кг/день\n",
			a.GetName(), inventoryItem.GetNumber(), animalType, a.GetFood())

		switch v := a.(type) {
		case *models.Tiger:
			fmt.Printf("  Сила: %d\n", v.Power)

		case *models.Wolf:
			fmt.Printf("  Ловкость: %d\n", v.Agility)

		case *models.Monkey:
			fmt.Printf("  Доброта: %d | Интеллект: %d\n", v.Kindness, v.Intelligence)

		case *models.Rabbit:
			fmt.Printf("  Доброта: %d\n", v.Kindness)
		}
	}
}
