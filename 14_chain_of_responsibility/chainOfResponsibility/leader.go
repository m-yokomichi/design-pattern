package chainOfResponsibility

import (
	"fmt"
)

type Leader struct {
	name string
	next Support
}

func (leader *Leader) SetNext(s Support) Support{
	leader.next = s
	return s
}

func (leader *Leader) Resolve(l Level) {
	if (l <= 2) {
		fmt.Println("resolve by leader")
		return 
	}

	leader.next.Resolve(l)
}

func NewLeader(name string) Support{
	var leader Support
	leader = &Leader{
		name: name,
	}
	return leader
}