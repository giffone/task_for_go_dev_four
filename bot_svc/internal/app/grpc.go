package app

import (
	"bot_svc/internal/service"
	"bot_svc/pkg/proto"

	"google.golang.org/grpc"
)

func newGrpc(svc *service.Service) *grpc.Server {
	server := grpc.NewServer()
	// register service
	proto.RegisterMessageServiceServer(server, svc)
	return server
}
