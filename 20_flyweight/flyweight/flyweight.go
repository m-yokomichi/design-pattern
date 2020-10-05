package flyweight

type FlyWeight struct {
	id    int
	count int
}

func (f *FlyWeight) AddCount() {
	f.count++
}

func (f *FlyWeight) GetId() {
	return f.id
}
