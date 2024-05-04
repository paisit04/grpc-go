package ports

import "github.com/paisit04/grpc-go/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
