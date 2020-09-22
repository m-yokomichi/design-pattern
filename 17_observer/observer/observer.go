package observer

type Observer interface {
	Update(*Generator)
}