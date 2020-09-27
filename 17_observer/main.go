package main

import (
	"./observer"
)

func main() {
	var degit observer.Observer
	degit = &observer.Degit{}

	var graph observer.Observer
	graph = &observer.Graph{}

	generator := &observer.NumberGenerator{
		Generator: &observer.Generator{},
	}
	generator.AddObserver(degit)
	generator.AddObserver(graph)

	generator.Execute()
}
