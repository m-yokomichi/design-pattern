package abstructFactory

type MizutakiFactory struct{}

func (_ *MizutakiFactory) GetSoup() Soup {
	return CreateChickenBonesSoup()
}

func (_ *MizutakiFactory) GetVegetables() Vegetables {
	var vegetables Vegetables
	vegetables.AddVegetable(CreateLeek())
	vegetables.AddVegetable(CreateCabbage())
	return vegetables

}
