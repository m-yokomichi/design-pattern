package main

import (
	factory "./factory"
)

func main() {
	p := factory.NewImagawasPrint()
	hanzai1 := p.CreateCuttable("芋")
	hanzai2 := p.CreateCuttable("木")

	p.CreateCutPrint(hanzai1)
	p.CreateCutPrint(hanzai2)
}
