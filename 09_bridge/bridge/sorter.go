package bridge

type Sortlmple interface {
	Sort()
}

type Sort struct {
	Sortlmple
}

func CreateSort(sortlmple Sortlmple) *Sort {
	sort := &Sort{
		Sortlmple : sortlmple,
	}

	return sort
}