package abstructFactory

type Cabbage struct {
	name string
}

func CreateCabbage() Vegetable {
	var cabbage Vegetable
	cabbage = Cabbage{"キャベツ"}
	return cabbage
}
