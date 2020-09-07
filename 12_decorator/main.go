package main

import (
	"fmt"

	"./decorator"
)

func main() {
	ice := decorator.CreateVanillaIcecream()
	fmt.Println(ice.GetName())
	fmt.Println(ice.HowSweet())

	toppingIce := decorator.CreateCasheNutsTopping(ice)
	fmt.Println(toppingIce.GetName())
	fmt.Println(toppingIce.HowSweet())
}
