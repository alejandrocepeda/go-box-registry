package boxes

import (
	"context"
	"fmt"
	"payment-rewrite/box"
)

// Self registration via Go package initialization
func init() {
	box.RegisterBox(func(deps box.Deps) box.Box {
		return StoreDB{db: deps.DB}
	})
}

type StoreDB struct {
	db box.DBPool
}

func (s StoreDB) ID() string { return "store_db" }

func (s StoreDB) Execute(ctx context.Context, input any) (box.Result, error) {
	data, ok := input.(map[string]any)
	if !ok {
		return box.Result{}, fmt.Errorf("store_db: invalid input type")
	}

	orderID, _ := data["order_id"].(string)

	// Read capture output from previous Box
	captureOutput, _ := data["capture"].(map[string]any)
	captureID, _ := captureOutput["capture_id"].(string)

	// Idempotent: INSERT ... ON CONFLICT DO NOTHING
	err := s.db.UpsertPaymentRecord(ctx, orderID, map[string]any{
		"capture_id": captureID,
	})
	if err != nil {
		return box.Result{}, fmt.Errorf("store_db: %w", err)
	}

	return box.Result{
		Output: map[string]any{
			"stored":     true,
			"capture_id": captureID,
		},
	}, nil
}
