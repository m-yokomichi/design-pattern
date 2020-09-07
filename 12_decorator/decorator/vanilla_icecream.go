package decorator

type VanillaIcecream struct {
	name string
	sweet string
}

func CreateVanillaIcecream() Icecream {
	var ice Icecream
	ice = &VanillaIcecream{
		name: "バニラアイス",
		sweet: "バニラ味",
	}
	return ice
}

func (i *VanillaIcecream) GetName() string {
	return i.name
}

func (i *VanillaIcecream) HowSweet() string {
	return i.sweet
}