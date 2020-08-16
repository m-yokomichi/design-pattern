package factory

type Print struct {}

func (p *Print) Draw(hanzai Cuttable) {}
func (p *Print) Cat(hanzai Cuttable) {}
func (p *Print) Prints(hanzai Cuttable) {}

func (p *Print) CreateCuttable() Cuttable{
	var wood = &Wood{}
	wood.Init()
	return wood
}

func (p *Print) CreateCutPrint(hanzai Cuttable) {
	p.Draw(hanzai)
	p.Cat(hanzai)
	p.Prints(hanzai)
}