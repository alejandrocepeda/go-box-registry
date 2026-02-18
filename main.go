package main

import (
	"fmt"
	"log"
	"payment-rewrite/box"
	_ "payment-rewrite/boxes"
	"payment-rewrite/deps"
)

func main() {
	boxRegistry := box.BuildAll()

	// The runtime mocks
	fmt.Printf("\n\n")
	fmt.Printf("Registered boxes: %v\n", boxRegistry.IDs())

	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	authorizeBox, err := boxRegistry.Get("authorize")
	if err != nil {
		log.Fatal(err)
	}

	authorizeResult, err := authorizeBox.Apply(box.NewDeps(
		deps.Job{ID: "JOB-001"},
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", authorizeBox.ID)
	fmt.Printf("Result: %+v\n", authorizeResult.Output)
	fmt.Printf("===============================================\n")

	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	captureBox, err := boxRegistry.Get("capture")
	if err != nil {
		log.Fatal(err)
	}

	captureResult, err := captureBox.Apply(box.NewDeps(
		deps.Job{ID: "JOB-001"},
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", captureBox.ID)
	fmt.Printf("Result: %+v\n", captureResult.Output)
	fmt.Printf("===============================================\n")
	fmt.Printf("\n")
	fmt.Printf("===============================================\n")
	storeDBBox, err := boxRegistry.Get("store_db")
	if err != nil {
		log.Fatal(err)
	}

	storeDBResult, err := storeDBBox.Apply(box.NewDeps(
		deps.DB{URL: "postgres://localhost/payments"},
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Box ID: %s\n", storeDBBox.ID)
	fmt.Printf("Result: %+v\n", storeDBResult.Output)
	fmt.Printf("===============================================\n")
}
