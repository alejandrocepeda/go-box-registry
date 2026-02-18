package main

import (
	"context"
	"fmt"
	"log"
	"payment-rewrite/box"
	_ "payment-rewrite/boxes"
)

type MockBankart struct{}

func (m MockBankart) Authorize(_ context.Context, orderID string, amount int64) (string, error) {
	return "auth_001", nil
}

func (m MockBankart) Capture(_ context.Context, orderID string, amount int64) (string, error) {
	return "cap_123", nil
}

type MockDB struct{}

func (m MockDB) UpsertPaymentRecord(_ context.Context, orderID string, data map[string]any) error {
	return nil
}

func main() {
	boxRegistry := box.BuildAll(box.Deps{
		Bankart: MockBankart{},
		DB:      MockDB{},
	})

	fmt.Printf("\n\n")
	fmt.Printf("Registered boxes: %v\n", boxRegistry.IDs())

	// Mocking the runtime
	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	authorizeBox, err := boxRegistry.Get("authorize")
	if err != nil {
		log.Fatal(err)
	}

	authorizeResult, err := authorizeBox.Execute(context.Background(), map[string]any{
		"order_id": "ORD-123",
		"amount":   int64(15000),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", authorizeBox.ID())
	fmt.Printf("Result: %+v\n", authorizeResult.Output)
	fmt.Printf("===============================================\n")
	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	storeDBBox, err := boxRegistry.Get("store_db")
	if err != nil {
		log.Fatal(err)
	}

	storeDBBoxResult, err := storeDBBox.Execute(context.Background(), map[string]any{
		"order_id": "ORD-123",
		"amount":   int64(15000),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", storeDBBox.ID())
	fmt.Printf("Result: %+v\n", storeDBBoxResult.Output)
	fmt.Printf("===============================================\n")
	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	captureBox, err := boxRegistry.Get("capture")
	if err != nil {
		log.Fatal(err)
	}

	captureResult, err := captureBox.Execute(context.Background(), map[string]any{
		"order_id": "ORD-123",
		"amount":   int64(15000),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", captureBox.ID())
	fmt.Printf("Result: %+v\n", captureResult.Output)
	fmt.Printf("===============================================\n")
	fmt.Printf("\n")

}
