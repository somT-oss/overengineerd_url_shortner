package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/somT-oss/urlshortner/internals/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
	port int
}

func NewAdapter(api ports.APIPort) (*Adapter, error) {
	return &Adapter{
		api: api,
	}, nil
}

func (adapter Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", adapter.port))
	if err != nil {
		log.Fatalf("failed to listen to server on port: %d", adapter.port)
	}
	grpcServer := grpc.NewServer()
	
	// Add config and register url shortening endpoints
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failted to serve the grpc server on port %d", adapter.port)
	}
}