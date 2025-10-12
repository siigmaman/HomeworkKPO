package models

type Computer struct {
	Thing
	CPU string
}

func NewComputer(name string, num int, cpu string) Computer {
	return Computer{Thing{name, num}, cpu}
}
