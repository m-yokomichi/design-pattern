package ingrediend

type Ingrediend struct {
	operandString string
}

func NewIngrediend(operandString string) Ingrediend {
	var operand Operand
	operand = &Ingrediend{
		operandString: operandString,
	}

	return operand
}

func (i *Ingrediend) GetOperandString() string {
	return i.operandString
}
