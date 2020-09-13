package chainOfResponsibility

import (
	"fmt"
)

type NoSupport struct {
	name string
	next Support
}

func (noSupport *NoSupport) SetNext(s Support) Support{
	noSupport.next = s
	return s
}

func (noSupport *NoSupport) Resolve(l Level) {
	fmt.Println("no resolve")
}