package observer

import (
	"math/rand"
	"time"
)

type NumberGenerator struct {
	*Generator
}

func (n *NumberGenerator) GetNumber() int {
	return n.number
}

func (n *NumberGenerator) Execute() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		n.number = rand.Intn(20)
		n.Generator.NotifyObservers()
	}
}
