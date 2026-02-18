package boxes

import (
	"context"
	"fmt"
	"payment-rewrite/box"
)

// Self registration via Go package initialization
func init() {
	box.RegisterBox(func(deps box.Deps) box.Box {
		return Capture{provider: deps.Bankart}
	})
}

type Capture struct {
	provider box.PaymentProvider
}

func (c Capture) ID() string { return "capture" }

func (c Capture) Execute(ctx context.Context, input any) (box.Result, error) {
	data, ok := input.(map[string]any)
	if !ok {
		return box.Result{}, fmt.Errorf("capture: invalid input type")
	}

	orderID, _ := data["order_id"].(string)
	amount, _ := data["amount"].(int64)

	captureID, err := c.provider.Capture(ctx, orderID, amount)
	if err != nil {
		return box.Result{}, fmt.Errorf("capture: %w", err)
	}

	return box.Result{
		Output: map[string]any{"capture_id": captureID},
	}, nil
}
