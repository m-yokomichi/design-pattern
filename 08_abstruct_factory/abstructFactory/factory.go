package abstructFactory

type Factory interface {
	GetSoup() Soup
	GetVegetables() Vegetables
}
