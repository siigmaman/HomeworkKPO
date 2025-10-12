package models

type Tiger struct {
	Predator
	Power int
}

func NewTiger(name string, food, power, number int) Tiger {
	return Tiger{
		Predator: Predator{
			Animal: Animal{Name: name, Food: food, Number: number},
		},
		Power: power,
	}
}
