package boxes

import (
	"context"
	"fmt"
	"payment-rewrite/box"
)

// Self registration via Go package initialization
func init() {
	box.RegisterBox(func(deps box.Deps) box.Box {
		return Authorize{provider: deps.Bankart}
	})
}

type Authorize struct {
	provider box.PaymentProvider
}

func (c Authorize) ID() string { return "authorize" }

func (c Authorize) Execute(ctx context.Context, input any) (box.Result, error) {
	data, ok := input.(map[string]any)
	if !ok {
		return box.Result{}, fmt.Errorf("authorize: invalid input type")
	}

	orderID, _ := data["order_id"].(string)
	amount, _ := data["amount"].(int64)

	authorizeID, err := c.provider.Authorize(ctx, orderID, amount)
	if err != nil {
		return box.Result{}, fmt.Errorf("authorize: %w", err)
	}

	return box.Result{
		Output: map[string]any{"authorize_id": authorizeID, "amount": amount},
	}, nil
}
