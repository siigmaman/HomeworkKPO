package models

type Animal struct {
	Name   string
	Food   int
	Number int
}

func (a Animal) GetFood() int      { return a.Food }
func (a *Animal) SetFood(food int) { a.Food = food }
func (a Animal) GetName() string   { return a.Name }
func (a Animal) GetNumber() int    { return a.Number }
