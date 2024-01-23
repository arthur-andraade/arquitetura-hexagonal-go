package dto

import "arquitetura-hexagonal/application"

type Product struct {
	ID     string  `json:"idProduct"`
	Name   string  `json:"nameProduct"`
	Price  float64 `json:"priceProduct"`
	Status string  `json:"productStatus"`
}

func NewProduct() *Product {
	return &Product{}
}

func (dto *Product) Bind(product *application.Product) (*application.Product, error) {

	if dto.ID != "" {
		product.ID = dto.ID
	}

	product.Name = dto.Name
	product.Price = dto.Price
	product.Status = dto.Status

	_, err := product.IsValid()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (dto *Product) Rebind(product application.ProductInterface) {
	dto.ID = product.GetID()
	dto.Name = product.GetName()
	dto.Price = product.GetPrice()
	dto.Status = product.GetStatus()
}
