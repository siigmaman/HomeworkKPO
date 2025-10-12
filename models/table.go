package models

type Table struct {
	Thing
	Material string
}

func NewTable(name string, num int, material string) Table {
	return Table{Thing{name, num}, material}
}
