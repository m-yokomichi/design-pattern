package main

import (
	factory "./factory"
)

func main() {
	p := factory.NewImagawasPrint()
	hanzai1 := p.CreateCuttable()
	p.CreateCutPrint(hanzai1)
}
