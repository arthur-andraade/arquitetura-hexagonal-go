package application

import (
	"errors"

	"github.com/google/uuid"
)

const (
	DISABLED = "DISABLED"
	ENABLED  = "ENABLED"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.New().String(),
		Status: DISABLED,
	}
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) IsValid() (bool, error) {

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be ENABLED or DISABLED")
	}

	priceIsValid, erro := p.validIfPriceIsGreaterOrEqualZero()

	if !priceIsValid {
		return priceIsValid, erro
	}

	return true, nil
}

func (p *Product) Enable() error {

	if _, erro := p.validIfPriceIsGreaterOrEqualZero(); erro != nil {
		return erro
	}

	p.Status = ENABLED

	return nil
}

func (p *Product) Disable() error {

	if _, erro := p.validIfPriceIsGreaterOrEqualZero(); erro != nil {
		return erro
	}

	p.Status = DISABLED

	return nil
}

func (p *Product) validIfPriceIsGreaterOrEqualZero() (bool, error) {

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	return true, nil
}
