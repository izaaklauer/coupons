package server

import (
	"context"
	"log"

        "github.com/izaaklauer/coupons/config"
	couponsv1 "github.com/izaaklauer/coupons/gen/proto/go/coupons/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CouponsServer struct {
	couponsv1.UnimplementedCouponsServiceServer

	config config.Coupons
}

// NewCouponsServer initializes a new server from config
func NewCouponsServer(config config.Coupons) (*CouponsServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := CouponsServer{
		config: config,
	}

	return &server, nil
}

func (s * CouponsServer) HelloWorld(
	ctx context.Context,
	req *couponsv1.HelloWorldRequest,
) (*couponsv1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &couponsv1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
