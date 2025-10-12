package di

import (
	"zoo/menu"
	"zoo/services"
)

type Container struct {
	Clinic *services.Clinic
	Zoo    *services.Zoo
	Menu   *menu.Menu
}

func NewContainer() *Container {
	clinic := services.NewClinic()
	zoo := services.NewZoo(clinic)
	menu := menu.NewMenu(zoo)
	return &Container{
		Clinic: clinic,
		Zoo:    zoo,
		Menu:   menu,
	}
}
