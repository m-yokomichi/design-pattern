package main

import (
	"fmt"

	singleton "./singleton"
)

func main() {
	s1 := singleton.GetSingletonStruct()
	s2 := singleton.GetSingletonStruct()

	fmt.Println(*s1)
	fmt.Println(*s2)
}
