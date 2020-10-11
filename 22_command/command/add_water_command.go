package command

import (
	"fmt"
)

type AddWaterCommand struct {
	*Command
}

func (c *AddWaterCommand) Execute() {
	for c.IsMelted() {
		c.AddWater(10)
		c.Mix()
	}

	fmt.Println("水を10gずつ加える実験")
	c.Note()
}