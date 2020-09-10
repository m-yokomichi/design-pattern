package visitor

type SuzukiHome struct {
	name string
}

func (h *SuzukiHome) Accept(t Teacher) {
	t.Visit(h)
}
