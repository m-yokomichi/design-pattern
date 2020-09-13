package chainOfResponsibility

import (
	"fmt"
)

type Member struct {
	name string
	next Support
}

func (m *Member) SetNext(r Support) Support{
	m.next = r
	return r
}

func (m *Member) Resolve(l Level) {
	if (l <= 1) {
		fmt.Println("resolve by member")
		return 
	}

	m.next.Resolve(l)
}

func NewMember(name string) Support{
	var member Support
	member = &Member{
		name: name,
	}
	return member
}