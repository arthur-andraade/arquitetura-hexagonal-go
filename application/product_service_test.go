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
