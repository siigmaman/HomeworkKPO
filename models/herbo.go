package models

type Herbo struct {
	Animal
	Kindness int
}

func (h Herbo) IsFriendly() bool {
	return h.Kindness > 5
}
