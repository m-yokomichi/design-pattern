package strategy

type AgeComparator struct{}

func (c *AgeComparator) Compare(h1, h2 Human) int64 {
	if (h1.Age < h2.Age) {
		return -1
	} else if (h1.Age == h2.Age) {
		return 0
	}
	return 1
}
