package application_test

import (
	"arquitetura-hexagonal/application"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_IsValid_mustReturnTrue_whenPriceIsGreaterOrEqualThanZero(t *testing.T) {
	product := productMock()
	isValid, err := product.IsValid()

	require.True(t, isValid)
	require.Nil(t, err)
}

func TestProduct_IsValid_mustReturnFalseAndErro_whenPriceIsSmallerThanZero(t *testing.T) {
	product := productMock()
	product.Price = -1
	isValid, err := product.IsValid()

	require.False(t, isValid)
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_IsValid_mustReturnFalseAndErro_whenStatusIsNotPermited(t *testing.T) {
	product := productMock()
	product.Status = "TESTE"
	isValid, err := product.IsValid()

	require.False(t, isValid)
	require.Equal(t, "the status must be ENABLED or DISABLED", err.Error())
}

func TestProduct_Enable_mustReturnError_whenPriceIsSmallerThanZero(t *testing.T) {

	product := productMock()
	product.Price = -1

	err := product.Enable()

	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_Disable_mustReturnError_whenPriceIsSmallerThanZero(t *testing.T) {

	product := productMock()
	product.Price = -1

	err := product.Disable()

	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_Enable_mustChangeStatusToEnabled_whenCallFunction(t *testing.T) {

	product := productMock()
	product.Status = application.DISABLED

	err := product.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)
}

func TestProduct_Disable_mustChangeStatusToDisabled_whenCallFunction(t *testing.T) {

	product := productMock()

	err := product.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
}

func productMock() application.Product {
	product := application.Product{
		ID:     "141",
		Name:   "Celular",
		Price:  10,
		Status: application.ENABLED,
	}

	return product
}
