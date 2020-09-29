package main

import (
	"fmt"

	"./memento"
)

func main() {
	taro := memento.Taro{}
	taro.SetInitMoney(10000)

	apple := memento.Item{"りんご", 100}
	orange := memento.Item{"みかん", 50}

	init := taro.CreateMemento()
	taro.Shopping(apple)
	taro.Shopping(apple)
	taro.Shopping(orange)

	fmt.Println(taro)
	fmt.Println(init)
}
