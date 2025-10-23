package payment

import "github.com/tiagorlampert/CHAOS/entities"

type PaymentService interface {
	CreatePayment(memeCoinID, userID string, amount float64, currency, paymentMethod string) (*entities.Payment, error)
	ProcessPayment(paymentID string) error
	GetPaymentByID(paymentID string) (*entities.Payment, error)
	GetPaymentsByUser(userID string) ([]entities.Payment, error)
	RefundPayment(paymentID string) error
	ValidatePaymentAmount(network string, amount float64) error
}