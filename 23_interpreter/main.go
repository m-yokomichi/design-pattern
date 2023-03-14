package main

import (
	"fmt"

	"./interpreter"
)

func main() {
	cap := interpreter.NewIngrediend("カップラーメン")
	hotwater := interpreter.NewIngrediend("お湯")
	expression := interpreter.NewExpression(interpreter.NewPlus(cap, hotwater))

	fmt.Println(expression.GetOperandString())
}
