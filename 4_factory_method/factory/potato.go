package factory

type Potato struct {
	materialName string
}

func (p *Potato) GetMaterialName() string {
	return p.materialName
}

func (p *Potato) Init() {
	p.materialName = "芋"

	return 
}