package chainOfResponsibility

import (
	"fmt"
)

type Manager struct {
	name string
	next Support
}

func (manager *Manager) SetNext(s Support) Support{
	manager.next = s
	return s
}

func (manager *Manager) Resolve(l Level) {
	if (l <= 3) {
		fmt.Println("resolve by manager")
		return 
	}

	manager.next.Resolve(l)
}

func NewManager(name string) Support{
	var manager Support
	manager = &Manager{
		name: name,
	}
	return manager
}