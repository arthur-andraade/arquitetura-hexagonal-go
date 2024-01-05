package application_test

import (
	"arquitetura-hexagonal/application"
	mock_application "arquitetura-hexagonal/application/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {

	controlador := gomock.NewController(t)

	productInterfaceMock := mock_application.NewMockProductInterface(controlador)

	t.Run("Must return product", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		persistenceMock.EXPECT().Get(gomock.Any()).Return(productInterfaceMock, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		productReturned, erro := service.Get("141")

		require.Nil(t, erro)
		require.Equal(t, productInterfaceMock, productReturned)

	})

	t.Run("Must return error when happens something wrong", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		persistenceMock.EXPECT().Get(gomock.Any()).Return(nil, errors.New("Erro")).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		productReturned, erro := service.Get("141")

		require.NotNil(t, erro)
		require.Nil(t, productReturned)
	})

	defer controlador.Finish()
}

func TestProductService_Create(t *testing.T) {

	controlador := gomock.NewController(t)

	productInterfaceMock := mock_application.NewMockProductInterface(controlador)

	t.Run("Price invalid", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		persistenceMock.EXPECT().Save(gomock.Any()).Times(0)

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		productReturned, err := service.Create("Laptop", -500)

		require.Nil(t, productReturned)
		require.NotNil(t, err)
		require.Equal(t, "the price must be greater or equal zero", err.Error())

	})

	t.Run("Return product created and saved", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		persistenceMock.EXPECT().Save(gomock.Any()).Return(productInterfaceMock, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		productReturned, err := service.Create("Laptop", 5000)

		require.Nil(t, err)
		require.NotNil(t, productReturned)

	})

	t.Run("Return erro when try save the new product", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		persistenceMock.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Erro")).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		productReturned, err := service.Create("Laptop", 5000)

		require.NotNil(t, err)
		require.Nil(t, productReturned)
	})

	defer controlador.Finish()
}
