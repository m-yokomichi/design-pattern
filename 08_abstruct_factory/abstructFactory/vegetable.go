package abstructFactory

type Vegetable interface {
}

type Vegetables struct {
	vegetables []Vegetable
}

func (vs *Vegetables) AddVegetable(v Vegetable) {
	vs.vegetables = append(vs.vegetables, v)
}
