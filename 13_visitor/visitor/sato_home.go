package visitor

type SatoHome struct {
	name string
}

func (h *SatoHome) Accept(t Teacher) {
	t.Visit(h)
}
