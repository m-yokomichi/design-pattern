package main

import (
	"fmt"

	"./command"
)

func main() {
	beaker1 := command.NewBeaker(10, 100)

	salt := command.AddSaltCommand{
		Command: &command.Command{},
	}
	salt.SetBeaker(beaker1)
	fmt.Println(salt.Command.Beaker)
	salt.Execute()

	beaker2 := command.NewBeaker(100, 100)
	water := command.AddWaterCommand{
		Command: &command.Command{},
	}
	water.SetBeaker(beaker2)
	water.Execute()
}
