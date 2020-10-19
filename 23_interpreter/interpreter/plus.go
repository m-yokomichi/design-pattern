package interpreter

type Plus struct {
	operand1 *Operand
	operand2 *Operand
}

func NewPlus(operand1, operand2 *Operand) *Plus {
	return &Plus{
		operand1: operand1,
		operand2: operand2,
	}
}

func (p *Plus) Execute() Operand {
	return NewIngrediend(p.operand1.GetOperandString() + "と" + p.operand2.GetOperandString() + "を足したもの")
}
