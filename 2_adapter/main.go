package main

import (
	adapter "2_adapter/adapter"
)

func main() {
	var newHello adapter.HelloInterface
	newHello = &adapter.NewHello{}
	newHello.SetString("World")

	// Interfaceに書かれていない機能は使えない　（Adapterパターンできない？）
	//newHello.CallHello()

	newHello.CallToHelloString()
}
