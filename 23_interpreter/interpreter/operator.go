package interpreter

type Operator interface {
	Execute() Operand
}
