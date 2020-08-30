package strategy

type WeightComparator struct{}

func (_ *WeightComparator) Compare(h1, h2 Human) int64 {
	if h1.Weight < h2.Weight {
		return -1
	} else if h1.Weight == h2.Weight {
		return 0
	}
	return 1
}
