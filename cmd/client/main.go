package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lgarron/multiplier/rpc/multiplier"
)

func main() {
	var A int32 = 4
	var B int32 = 5

	client := multiplier.NewMultiplierProtobufClient(("http://localhost:9000"), &http.Client{})
	product, err := client.Multiply(context.Background(), &multiplier.Arguments{A: A, B: B})
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}
	value := product.GetValue()
	fmt.Printf("Product of %d and %d is: %d", A, B, value)
}
