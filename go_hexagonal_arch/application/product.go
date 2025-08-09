package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float32
	Status string
}

// func (product *Product) IsValid() (bool, error) {

// }

func (product *Product) Enable() error {
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	product.Status = ENABLED

	return nil
}

// func (product *Product) Disable() error {

// }

func (product *Product) GetID() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetStatus() string {
	return product.Status
}

func (product *Product) GetPrice() float32 {
	return product.Price
}
