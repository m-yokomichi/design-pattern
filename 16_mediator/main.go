package main

import(
	"./mediator"
	"fmt"
)

func main() {
	user := mediator.CreateUser()
	fmt.Println(user.CheckTweet())
}
