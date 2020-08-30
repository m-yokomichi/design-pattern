package strategy

type Comparator interface {
	Compare(Human, Human) int64
}