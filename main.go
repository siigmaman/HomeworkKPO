package main

import (
	"zoo/di"
)

func main() {
	c := di.NewContainer()
	c.Menu.Run()
}
