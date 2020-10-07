package proxy

import (
	"fmt"
)

type Principal struct {
	Name string
}

func (p *Principal) Method1() {
	fmt.Println(p.Name, "のMethod1")
}

func (p *Principal) Method2() {
	fmt.Println(p.Name, "のMethod2")
}
