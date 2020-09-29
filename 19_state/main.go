package main

import (
	"fmt"

	"./state"
)

func main() {
	var goodState state.State
	goodState = &state.GoodState{}

	var badState state.State
	badState = &state.BadState{}

	taro := &state.Taro{}
	taro.ChangeState(goodState)
	fmt.Println(taro.ShowTension())

	taro.ChangeState(badState)
	fmt.Println(taro.ShowTension())
}
