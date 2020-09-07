package decorator

type CasheNutsTopping struct {
	ice  Icecream
	name string
}

func CreateCasheNutsTopping(baseIce Icecream) Icecream {
	var newIce Icecream
	newIce = &CasheNutsTopping{
		ice:  baseIce,
		name: "カシューナッツ",
	}

	return newIce
}
func (c *CasheNutsTopping) GetName() string {
	return c.name + c.ice.GetName()
}

func (c *CasheNutsTopping) HowSweet() string {
	return c.ice.HowSweet()
}
