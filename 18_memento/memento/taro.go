package memento

type Taro struct {
	money int
	items []Item
}

func (t *Taro) SetInitMoney(money int) {
	t.money = money
}

func (t *Taro) GetMoney() int {
	return t.money
}

func (t *Taro) Shopping(item Item) {
	t.money -= item.Price
	t.items = append(t.items, item)
}

func (t *Taro) CreateMemento() Memento {
	return Memento {
		Money : t.money,
		Items : t.items,
	}
}