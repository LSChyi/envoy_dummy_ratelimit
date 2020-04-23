package grpc

import (
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type service interface {
	RegisterFn(*grpc.Server)
}

type server struct {
	listenAddr string
	*grpc.Server
}

func NewGRPCServer(listenAddr string) *server {
	return &server{
		listenAddr: listenAddr,
		Server:     grpc.NewServer(),
	}
}

func (s *server) Register(srv service) {
	srv.RegisterFn(s.Server)
}

func (s *server) Run() {
	lis, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.WithField("address", s.listenAddr).Info("starts listening on address")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
