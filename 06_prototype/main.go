package main

import (
	"fmt"

	"./prototype"
)

func main() {
	base := prototype.CreateMold("星")
	p := base.CreateClone()

	fmt.Println(base)
	fmt.Println(p)
}
