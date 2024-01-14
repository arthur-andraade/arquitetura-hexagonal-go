package cli_test

import (
	"arquitetura-hexagonal/adapters/cli"
	"arquitetura-hexagonal/application"
	mock_application "arquitetura-hexagonal/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc123"
	productName := "Playstation 4 Slim"
	productPrice := 2500.0

	productReturnMock := mock_application.NewMockProductInterface(ctrl)
	productReturnMock.EXPECT().GetID().Return(productId).AnyTimes()
	productReturnMock.EXPECT().GetName().Return(productName).AnyTimes()
	productReturnMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productReturnMock.EXPECT().GetStatus().Return(application.ENABLED).AnyTimes()

	t.Run("Action create", func(t *testing.T) {
		productReturnMock.EXPECT().GetStatus().Return(application.ENABLED).AnyTimes()
		service := mock_application.NewMockProductServiceInterface(ctrl)
		service.EXPECT().Get(productId).Return(productReturnMock, nil).AnyTimes()
		service.EXPECT().Create(productName, productPrice).Return(productReturnMock, nil).AnyTimes()

		result, err := cli.Run(service, "create", productId, productName, productPrice)

		require.Nil(t, err)
		require.Equal(t, "Product ID abc123 with the name Playstation 4 Slim has been created with the price 2500.00 and status ENABLED", result)
	})

	t.Run("Action enable", func(t *testing.T) {
		service := mock_application.NewMockProductServiceInterface(ctrl)
		service.EXPECT().Get(productId).Return(productReturnMock, nil).AnyTimes()
		service.EXPECT().Enable(gomock.Any()).Return(productReturnMock, nil).AnyTimes()

		result, err := cli.Run(service, "enable", productId, "", 0.0)

		require.Nil(t, err)
		require.Equal(t, "Product ID abc123 with the name Playstation 4 Slim has been changed the status to ENABLED", result)
	})

	t.Run("Action disable", func(t *testing.T) {
		service := mock_application.NewMockProductServiceInterface(ctrl)
		service.EXPECT().Get(productId).Return(productReturnMock, nil).AnyTimes()
		service.EXPECT().Disable(gomock.Any()).Return(productReturnMock, nil).AnyTimes()

		result, err := cli.Run(service, "disable", productId, "", 0.0)

		require.Nil(t, err)
		require.Equal(t, "Product ID abc123 with the name Playstation 4 Slim has been changed the status to ENABLED", result)
	})

	t.Run("Action default that return product saved", func(t *testing.T) {
		service := mock_application.NewMockProductServiceInterface(ctrl)
		service.EXPECT().Get(productId).Return(productReturnMock, nil).AnyTimes()

		result, err := cli.Run(service, "", productId, "", 0.0)

		require.Nil(t, err)
		require.Equal(t, "Product ID abc123 | Name Playstation 4 Slim | Price 2500.00 | Status ENABLED", result)
	})

}
