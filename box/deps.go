package box

import "context"

type PaymentProvider interface {
	Authorize(ctx context.Context, orderID string, amount int64) (string, error)
	Capture(ctx context.Context, orderID string, amount int64) (string, error)
}

type DBPool interface {
	UpsertPaymentRecord(ctx context.Context, orderID string, data map[string]any) error
}

type LogisticsClient interface {
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

type EcommerceClient interface {
	UpdateStatus(ctx context.Context, orderID string, status string) error
}

type Deps struct {
	Bankart   PaymentProvider
	DB        DBPool
	Logistics LogisticsClient
	Ecommerce EcommerceClient
}
