package services

import (
	"fmt"
	"zoo/asmutils"
	"zoo/models"
)

type IClinic interface {
	CheckHealth() bool
}

type Zoo struct {
	Animals []models.IAlive
	Things  []models.IInventory
	Clinic  IClinic
}

func NewZoo(clinic IClinic) *Zoo {
	return &Zoo{Clinic: clinic}
}

func (z *Zoo) AddAnimal(a models.IAlive) {
	if z.Clinic.CheckHealth() {
		z.Animals = append(z.Animals, a)
		fmt.Printf("Принято: %s\n", a.GetName())
	} else {
		fmt.Printf("Отклонено (болен): %s\n", a.GetName())
	}
	temp := CheckTemp()
	fmt.Printf("Температура при приёме: %.1f°C\n", temp)
}

func (z *Zoo) AddThing(t models.IInventory) {
	z.Things = append(z.Things, t)
}

func (z *Zoo) Report() {
	totalFood := 0
	for _, a := range z.Animals {
		totalFood = asmutils.SumFood(totalFood, a.GetFood())
	}

	fmt.Println("\nФормирование отчёта...")
	asmutils.SleepASM(50000)
	fmt.Printf("Готово!\n")

	fmt.Printf("\nВсего животных: %d\n", len(z.Animals))
	fmt.Printf("Общее потребление еды: %d кг/день\n\n", totalFood)

	if len(z.Animals) > 0 {
		avg := asmutils.CalcDailyFeedAverage(totalFood, len(z.Animals))
		fmt.Printf("Среднее потребление на одно животное: %d кг/день\n", avg)
	}

	fmt.Println("Контактный зоопарк:")
	for _, a := range z.Animals {
		if f, ok := a.(interface{ IsFriendly() bool }); ok && f.IsFriendly() {
			fmt.Println(" -", a.GetName())
		}
	}

	fmt.Println("\nИнвентарь:")
	for _, t := range z.Things {
		fmt.Printf(" - %v (№%d)\n", t, t.GetNumber())
	}
}
