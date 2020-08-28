package abstructFactory

type HotPot struct {
	soup       Soup
	vegetables Vegetables
}

func (h *HotPot) AddSoup(soup Soup) {
	h.soup = soup
}

func (h *HotPot) AddVegetables(vegetables Vegetables) {
	h.vegetables = vegetables
}
