package factory

import "fmt"

type ImagawasPrint struct {
	*Print
}

func NewImagawasPrint() *ImagawasPrint {
	imagawasPrint := &ImagawasPrint{
		Print: &Print{},
	}
	imagawasPrint.print = imagawasPrint
	return imagawasPrint
}

func (_ *ImagawasPrint) Draw(hanzai Cuttable) {
	fmt.Println(hanzai.GetMaterialName(), "に絵を描く")
}

func (_ *ImagawasPrint) Cat(hanzai Cuttable) {
	fmt.Println(hanzai.GetMaterialName(), "を切る")
}
func (_ *ImagawasPrint) Prints(hanzai Cuttable) {
	fmt.Println("プリントする")
}

func (_ *ImagawasPrint) CreateCuttable() *Cuttable {
	var hanzai Cuttable
	hanzai = NewPotato()
	return &hanzai
}

func (i *ImagawasPrint) CreateCutPrint(hanzai Cuttable) {
	i.Draw(hanzai)
	i.Cat(hanzai)
	i.Prints(hanzai)
}
