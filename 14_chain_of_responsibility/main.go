package main

import (
	c "./chainOfResponsibility" 
	"fmt"
)
func main() {
	member := c.NewMember("Jone")
	leader := c.NewLeader("Bob")
	manager := c.NewManager("Jully")
	noSupport := &c.NoSupport{}

	member.SetNext(leader).SetNext(manager).SetNext(noSupport)

	var level c.Level
	for level = 1; level <= 4; level++ {
		fmt.Println("level", level)
		member.Resolve(level)
	}
}