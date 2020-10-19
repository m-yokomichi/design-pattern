package interpreter

type Wait struct {
	minutes int
	operand Operand
}

func NewWait(minutes int, operand Operand) *Wait {
	return &Wait{
		minutes: minutes,
		operand: operand,
	}
}

func (w *Wait) Execute() Operand {
	return NewIngrediend(w.operand.GetOperandString() + "を" + string(w.minutes) + "置いたもの")
}
