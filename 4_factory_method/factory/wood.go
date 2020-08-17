package factory

type Wood struct {
	materialName string
}

func (w *Wood) GetMaterialName() string {
	return w.materialName
}
func NewWood() *Wood {
	wood := Wood{}
	wood.materialName = "芋"

	return &wood
}
