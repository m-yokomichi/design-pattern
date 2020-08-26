package abstructFactory

type Cabbage struct {
	name string
}

func CreateCabbage() Vegitable {
	var cabbage Vegitable
	cabbage = Cabbage{"キャベツ"}
	return cabbage
}