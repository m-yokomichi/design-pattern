package memento

type Memento struct {
	Money int
	Items []Item
}

func (m *Memento) GetMoney() int {
	return m.Money
}
func (m *Memento) GetItems() []Item {
	return m.Items
}
