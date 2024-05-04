package grpc

import (
	"context"

	"github.com/paisit04/grpc-go-proto/golang/payment"
	"github.com/paisit04/grpc-go/payment/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, err
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
