package abstructFactory

type ChickenBonesSoup struct {
	name string
}

func CreateChickenBonesSoup() Soup {
	var soup Soup
	soup = ChickenBonesSoup{"チキン"}
	return soup
}
