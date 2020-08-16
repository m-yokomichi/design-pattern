package factory

import (
	"fmt"
)

type ImagawasPrint struct {
	*Print
}

func (_ *ImagawasPrint) Draw(hanzai Cuttable) {
	fmt.Println(hanzai.GetMaterialName(), "に絵を描く")
}

func (_ *ImagawasPrint) Cat(hanzai Cuttable) {
	fmt.Println(hanzai.GetMaterialName(), "を切る")
}
func (_ *ImagawasPrint) Prints(hanzai Cuttable) {
	fmt.Println("プリントする")}

func (_ *ImagawasPrint) CreateCuttable() Cuttable {
	potato := &Potato {}
	potato.Init()
	return potato
}

func (i *ImagawasPrint) CreateCutPrint(hanzai Cuttable) {
	i.Draw(hanzai)
	i.Cat(hanzai)
	i.Prints(hanzai)
}