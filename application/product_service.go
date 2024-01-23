package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {

	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {

	newProduct := NewProduct()
	newProduct.Name = name
	newProduct.Price = price

	_, err := newProduct.IsValid()

	if err != nil {
		return nil, err
	}

	productSaved, err := s.Persistence.Save(newProduct)

	if err != nil {
		return nil, err
	}

	return productSaved, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {

	err := product.Enable()

	if err != nil {
		return product, err
	}

	return s.updateStatusProduct(product)
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {

	err := product.Disable()

	if err != nil {
		return product, err
	}

	return s.updateStatusProduct(product)
}

func (s *ProductService) updateStatusProduct(product ProductInterface) (ProductInterface, error) {

	_, err := s.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
