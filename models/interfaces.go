package models

type IAlive interface {
	GetFood() int
	SetFood(int)
	GetName() string
}

type IInventory interface {
	GetNumber() int
}
