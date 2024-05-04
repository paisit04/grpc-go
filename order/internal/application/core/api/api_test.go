package api

import (
	"errors"
	"testing"

	"github.com/paisit04/grpc-go/order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockedPayment struct {
	mock.Mock
}

func (p *mockedPayment) Charge(order *domain.Order) error {
	args := p.Called(order)
	return args.Error(0)
}

type mockedDb struct {
	mock.Mock
}

func (d *mockedDb) Save(order *domain.Order) error {
	args := d.Called(order)
	return args.Error(0)
}

func (d *mockedDb) Get(id string) (domain.Order, error) {
	args := d.Called(id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func Test_Should_Place_Order(t *testing.T) {
	payment := new(mockedPayment)
	payment.On("Charge", mock.Anything).Return(nil)

	db := new(mockedDb)
	db.On("Save", mock.Anything).Return(nil)

	application := NewApplication(db, payment)

	_, err := application.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.Nil(t, err)
}

func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mockedPayment)
	payment.On("Charge", mock.Anything).Return(nil)

	db := new(mockedDb)
	db.On("Save", mock.Anything).Return(errors.New("connection error"))

	application := NewApplication(db, payment)

	_, err := application.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "phone",
				UnitPrice:   14.7,
				Quantity:    1,
			},
		},
		CreatedAt: 0,
	})
	assert.EqualError(t, err, "connection error")
}

func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
	payment := new(mockedPayment)
	payment.On("Charge", mock.Anything).Return(errors.New("insufficient balance"))

	db := new(mockedDb)
	db.On("Save", mock.Anything).Return(nil)

	application := NewApplication(db, payment)

	_, err := application.PlaceOrder(domain.Order{
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "bag",
				UnitPrice:   2.5,
				Quantity:    6,
			},
		},
		CreatedAt: 0,
	})
	st, _ := status.FromError(err)
	assert.Equal(t, st.Message(), "order creation failed")
	assert.Equal(t, st.Details()[0].(*errdetails.BadRequest).FieldViolations[0].Description, "insufficient balance")
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}
