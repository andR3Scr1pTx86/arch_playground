package application_test

import (
	"testing"

	"github.com/andR3Scr1pTx86/arch_playground/go_hexagonal_arch/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "price must be greater than 0", err.Error())
}
