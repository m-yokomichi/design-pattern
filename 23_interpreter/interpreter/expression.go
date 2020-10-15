package interpreter

type Expression struct {
	operator *Operator
}

func NewExpression(operator *Operator) {
	var operand Operand
	operand = &Expression{
		operator: operator,
	}

	return operand
}

func (e *Expression) GetOperandString() string {
	return e.operator.execute().GetOperandString()
}
