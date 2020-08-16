package factory

type Wood struct {
	materialName string
}

func (w *Wood) GetMaterialName() string {
	return w.materialName
}

func (w *Wood) Init() {
	w.materialName = "木"

	return 
}