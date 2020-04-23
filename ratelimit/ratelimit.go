package ratelimit

import (
	"context"
	"fmt"

	pb "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v2"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	DestinationClusterKey         = "destination_cluster"
	ServerRequestDescriptorPrefix = "inbound"
)

type server struct{}

func NewRateLimitServer() *server {
	return new(server)
}

func (s *server) ShouldRateLimit(ctx context.Context, request *pb.RateLimitRequest) (*pb.RateLimitResponse, error) {
	log.WithField("request", request).Info("log out request content")
	if request == nil {
		return nil, fmt.Errorf("get nil request")
	}
	return new(pb.RateLimitResponse), nil
}

func (s *server) RegisterFn(grpcServer *grpc.Server) {
	pb.RegisterRateLimitServiceServer(grpcServer, s)
}
