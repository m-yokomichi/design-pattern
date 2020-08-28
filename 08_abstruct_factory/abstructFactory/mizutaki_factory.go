package abstructFactory

type MizutakiFactory struct{}

func (_ *MizutakiFactory) GetSoup() Soup {
	soup := CreateChickenBonesSoup()
	return *soup
}

func (_ *MizutakiFactory) GetVegetables() Vegetables {
	vegetables := Vegetables{}
	vegetables.AddVegetable(CreateLeek())
	vegetables.AddVegetable(CreateCabbage())
	return vegetables
}
