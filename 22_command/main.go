package main

import (
	"fmt"

	"./command"
)

func main() {
	beaker1 := command.NewBeaker(10, 100)
	fmt.Println(beaker1)
}
