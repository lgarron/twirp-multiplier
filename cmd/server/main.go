package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lgarron/multiplier/rpc/multiplier"
)

// Server impl
type Server struct{}

// Multiply impl
func (s *Server) Multiply(ctx context.Context, arguments *multiplier.Arguments) (product *multiplier.Product, err error) {
	return &multiplier.Product{
		Value: arguments.GetA() * arguments.GetB(),
	}, nil
}

func main() {
	server := &Server{}
	twirpHandler := multiplier.NewMultiplierServer(server, nil)

	err := http.ListenAndServe(":9000", twirpHandler)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hi!")
}
