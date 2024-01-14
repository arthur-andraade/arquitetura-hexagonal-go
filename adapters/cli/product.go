package cli

import (
	"arquitetura-hexagonal/application"
	"fmt"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {

	case "create":

		product, err := service.Create(productName, price)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID %s with the name %s has been created with the price %.2f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	case "enable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		product, err = service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID %s with the name %s has been changed the status to %s",
			product.GetID(),
			product.GetName(),
			product.GetStatus(),
		)

	case "disable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		product, err = service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID %s with the name %s has been changed the status to %s",
			product.GetID(),
			product.GetName(),
			product.GetStatus(),
		)

	default:

		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID %s | Name %s | Price %.2f | Status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus(),
		)

	}

	return result, nil
}
