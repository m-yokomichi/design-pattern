package factory

type Potato struct {
	materialName string
}

func (p *Potato) GetMaterialName() string {
	return p.materialName
}
func NewPotato() *Potato {
	potato := Potato{}
	potato.materialName = "芋"

	return &potato
}
