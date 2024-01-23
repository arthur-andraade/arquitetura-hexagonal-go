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

func TestProductService_Enable(t *testing.T) {

	controlador := gomock.NewController(t)

	t.Run("Enable the product and save it", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Enable().Return(nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Return(productInterfaceMock, nil).AnyTimes()

		_, err := service.Enable(productInterfaceMock)

		require.Nil(t, err)

	})

	t.Run("Enable the product but happen error to save it", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Enable().Return(nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Erro")).AnyTimes()

		_, err := service.Enable(productInterfaceMock)

		require.NotNil(t, err)

	})

	t.Run("Happen error when try enable the product", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Enable().Return(errors.New("Erro")).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Times(0)

		_, err := service.Enable(productInterfaceMock)

		require.NotNil(t, err)

	})

	defer controlador.Finish()

}

func TestProductService_Disable(t *testing.T) {

	controlador := gomock.NewController(t)

	t.Run("Disable the product and save it", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Disable().Return(nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Return(productInterfaceMock, nil).AnyTimes()

		_, err := service.Disable(productInterfaceMock)

		require.Nil(t, err)

	})

	t.Run("Disable the product but happen error to save it", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Disable().Return(nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Return(nil, errors.New("Erro")).AnyTimes()

		_, err := service.Disable(productInterfaceMock)

		require.NotNil(t, err)

	})

	t.Run("Happen error when try disable the product", func(t *testing.T) {

		persistenceMock := mock_application.NewMockProductPersistenceInterface(controlador)

		productInterfaceMock := mock_application.NewMockProductInterface(controlador)

		productInterfaceMock.EXPECT().Disable().Return(errors.New("Erro")).AnyTimes()

		service := application.ProductService{
			Persistence: persistenceMock,
		}

		persistenceMock.EXPECT().Save(gomock.Any()).Times(0)

		_, err := service.Disable(productInterfaceMock)

		require.NotNil(t, err)

	})

	defer controlador.Finish()

}
