package ports

import (
	"context"

	"github.com/paisit04/grpc-go/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
