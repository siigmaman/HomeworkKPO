package models

type Monkey struct {
	Herbo
	Intelligence int
}

func NewMonkey(name string, food, kindness, intelligence, number int) Monkey {
	return Monkey{
		Herbo: Herbo{
			Animal:   Animal{Name: name, Food: food, Number: number},
			Kindness: kindness,
		},
		Intelligence: intelligence,
	}
}
