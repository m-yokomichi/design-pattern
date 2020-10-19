package main

import (
	"./interpreter"
	"fmt"
)

func main() {
	cap := interpreter.NewIngrediend("カップラーメン")
	hotwater := interpreter.NewIngrediend("お湯")
	expression := interpreter.NewExpression(NewPlus(cap, hotwater))

	fmt.Println(expression.GetOperandString())
}