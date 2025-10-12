package models

type Thing struct {
	Name   string
	Number int
}

func (t Thing) GetNumber() int { return t.Number }
