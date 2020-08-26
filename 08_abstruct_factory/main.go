package main

import (
	"fmt"

	"./abstructFactory"
)

func main() {
	cabbage := abstructFactory.CreateCabbage()
	fmt.Println(cabbage)
}
