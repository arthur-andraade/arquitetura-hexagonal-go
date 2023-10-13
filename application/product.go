package application

import "errors"

const (
	DISABLED = "DISABLED"
	ENABLED  = "ENABLED"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
}

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
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
