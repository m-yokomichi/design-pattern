package main

import (
	"fmt"

	"./flyweight"
)

func main() {
	factory := flyweight.GetInstance()
	fmt.Println(factory)

	factory.GetFlyWeight(10)
	factory.GetFlyWeight(11)
	fmt.Println(factory)
}
