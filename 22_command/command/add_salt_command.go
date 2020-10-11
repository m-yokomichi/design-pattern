package command

import (
	"fmt"
)

type AddSaltCommand struct {
	*Command
}

func (c *AddSaltCommand) Execute() {
	for !c.IsMelted() {
		c.AddSalt(1)
		c.Mix()
		c.Note()
	}

	fmt.Println("塩を1gずつ加える実験")
	c.Note()
}