package models

type Rabbit struct {
	Herbo
}

func NewRabbit(name string, food, kindness, number int) Rabbit {
	return Rabbit{
		Herbo: Herbo{
			Animal:   Animal{Name: name, Food: food, Number: number},
			Kindness: kindness,
		},
	}
}
