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
	fmt.Println(hanzai.GetProductName(), "に絵を描く")
}

func (_ *ImagawasPrint) Cat(hanzai Cuttable) {
	fmt.Println(hanzai.GetProductName(), "を切る")
}
func (_ *ImagawasPrint) Prints(hanzai Cuttable) {
	fmt.Println("プリントする")
}

func (_ *ImagawasPrint) CreateCuttable(productName string) Cuttable {
	var hanzai Cuttable
	hanzai = CreateProduct(productName)
	return hanzai
}

func (i *ImagawasPrint) CreateCutPrint(hanzai Cuttable) {
	i.print.Draw(hanzai)
	i.print.Cat(hanzai)
	i.print.Prints(hanzai)
}
