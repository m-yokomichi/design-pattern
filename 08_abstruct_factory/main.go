package main

import (
	"fmt"

	"./abstructFactory"
)

func main() {
	var factory abstructFactory.Factory
	hotpot := abstructFactory.HotPot{}
	factory = abstructFactory.MizutakiFactory{}
	hotpot.AddSoup(factory.GetSoup())
	hotpot.AddVegetables(factory.GetVegetables())
	fmt.Println(hotpot)
}
