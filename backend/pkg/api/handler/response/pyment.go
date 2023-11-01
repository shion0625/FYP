package response

import "github.com/shion0625/backend/pkg/domain"

type OrderPayment struct {
	PaymentType  domain.PaymentType `json:"payment_type"`
	PaymentOrder any                `json:"payment_order"`
}
