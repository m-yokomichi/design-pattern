package interpreter

type Expression struct {
	operator Operator
}

func NewExpression(operator Operator) Operand {
	var operand Operand
	operand = &Expression{
		operator: operator,
	}

	return operand
}

func (e *Expression) GetOperandString() string {
	return e.operator.Execute().GetOperandString()
}
