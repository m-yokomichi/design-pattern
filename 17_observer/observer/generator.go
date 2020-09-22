package observer

type Generator struct {
	observers []Observer
}

func (g *Generator) AddObserver(observer Observer) {
	g.observers = append(g.observers, observer)
}
func (g *Generator) NotifyObservers() {
	for _, observer := range g.observers {
		observer.Update(g)
	}
}
