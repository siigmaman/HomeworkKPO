package models

type Wolf struct {
	Predator
	Agility int
}

func NewWolf(name string, food, agility, number int) Wolf {
	return Wolf{
		Predator: Predator{
			Animal: Animal{Name: name, Food: food, Number: number},
		},
		Agility: agility,
	}
}
