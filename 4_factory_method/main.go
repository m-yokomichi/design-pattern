package main

import (
	factory "4_factory_method/factory"
)

func main() {
	p := new(factory.Print)
	hanzai1 := p.CreateCuttable()
	p.CreateCutPrint(hanzai1)

	ip := new(factory.ImagawasPrint)
	hanzai2 := ip.CreateCuttable()
	ip.CreateCutPrint(hanzai2)
}