package main

import(
	"./facade"
	"fmt"
)

func main() {
	user := facade.User{}
	fmt.Println(user.GetUser())
}