package factory

type Product struct {
	productName string
}

func (p *Product) GetProductName() string {
	return p.productName
}
func CreateProduct(productName string) *Product {
	product := Product{}
	product.productName = productName

	return &product
}
