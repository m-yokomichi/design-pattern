package abstructFactory

type Leek struct {
	name string
}

func CreateLeek() Vegetable {
	var leek Vegetable
	leek = Leek{"ネギ"}
	return leek
}
