package observer

import (
	"fmt"
)
type Graph struct {}

func (g *Graph) Update(generator *Generator) {
	var graph string
	for i := 0; i < generator.GetNumber(); i++ {
		graph += "+"
	}

	fmt.Println(graph)
}