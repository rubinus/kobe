package main

import (
	"google.golang.org/grpc"
	"github.com/KubeOperator/kobe/api"
	"github.com/KubeOperator/kobe/pkg/server"
	"net"
)

func newTcpListener(address string) (*net.Listener, error) {
	s, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func newServer() *grpc.Server {
	gs := grpc.NewServer()
	kobe := server.NewKobe()
	api.RegisterKobeApiServer(gs, kobe)
	return gs
}
