package flyweight

type FlyWeight struct {
	id int
}

func (f *FlyWeight) GetId() int {
	return f.id
}
