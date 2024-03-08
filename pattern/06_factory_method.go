package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import (
	"errors"
	"fmt"
)

// PaymentGatewayType это тип Gateway
type PaymentGatewayType int

const (
	PayPalGatewayType PaymentGatewayType = iota
	StripeGatewayType
)

// PaymentGateway это общий интерфейс для всех Gateway
type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

type PayPalGateway struct{}

func (pg *PayPalGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}

type StripeGateway struct{}

func (sg *StripeGateway) ProcessPayment(amount float64) error {
	fmt.Printf("Processing Stripe payment of $%.2f\n", amount)
	return nil
}

// метод NewPaymentGateway создаёт структуру реализующую интерфейс PaymentGateway
// основываясь на типе.
// Это фабричный метод
func NewPaymentGateway(gwType PaymentGatewayType) (PaymentGateway, error) {
	switch gwType {
	case PayPalGatewayType:
		// здесь может любой конструктор
		return &PayPalGateway{}, nil
	case StripeGatewayType:
		return &StripeGateway{}, nil
	default:
		return nil, errors.New("unsupported payment gateway type")
	}
}
