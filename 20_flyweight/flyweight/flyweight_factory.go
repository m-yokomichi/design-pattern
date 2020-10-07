package flyweight

type FlyWeightFactory struct {
	pool map[int]FlyWeight
}

type FlyWeightFactoryInterface interface {
	GetFlyWeight(int) FlyWeight
}

var singleton FlyWeightFactoryInterface

func GetInstance() FlyWeightFactoryInterface {
	if singleton != nil {
		return singleton
	}

	singleton = &FlyWeightFactory{
		pool: map[int]FlyWeight{},
	}
	return singleton
}

func (f *FlyWeightFactory) GetFlyWeight(id int) FlyWeight {
	_, ok := f.pool[id]
	if ok == true {
		return f.pool[id]
	}

	f.pool[id] = FlyWeight{id: id}
	return f.pool[id]
}
