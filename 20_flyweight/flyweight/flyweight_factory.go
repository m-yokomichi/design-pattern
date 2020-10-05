package flyweight

type FlyWeightFactory struct {
	pool []FlyWeight
}

type singleton FlyWeightFactory

func GetInstance() {
	if singleton != nil {
		return &singleton
	}
	singleton = FlyWeightFactory{}
	return &singleton
}

func GetFlyWeight() {

}
