package proxy

import (
	"fmt"
)

type Proxy struct {
	*Principal
	Name string
}

func (p *Proxy) Method1() {
	fmt.Println(p.Name, "のMethod1")
}

func (p *Proxy) Method2() {
	p.Principal.Method2()
}
