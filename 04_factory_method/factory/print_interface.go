package factory

type PrintInterface interface {
	Draw(hanzai Cuttable)
	Cat(hanzai Cuttable)
	Prints(hanzai Cuttable)
}
type Print struct {
	print PrintInterface
}

func (p *Print) CreateCutPrint(hanzai Cuttable) {
	p.print.Draw(hanzai)
	p.print.Cat(hanzai)
	p.print.Prints(hanzai)
}
