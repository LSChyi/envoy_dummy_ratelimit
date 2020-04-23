package ratelimit

import (
	"context"
	"testing"

	ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
	pb "github.com/envoyproxy/go-control-plane/envoy/service/ratelimit/v2"
	"github.com/stretchr/testify/assert"
)

func TestEmptyShouldRateLimit(t *testing.T) {
	s := NewRateLimitServer()
	res, err := s.ShouldRateLimit(context.Background(), nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestSimpleShouldRateLimit(t *testing.T) {
	s := NewRateLimitServer()
	req := pb.RateLimitRequest{
		Domain: "bonobo",
		Descriptors: []*ratelimit.RateLimitDescriptor{
			&ratelimit.RateLimitDescriptor{
				Entries: []*ratelimit.RateLimitDescriptor_Entry{
					&ratelimit.RateLimitDescriptor_Entry{
						Key:   "test_key",
						Value: "test_value",
					},
				},
			},
		},
	}

	res, err := s.ShouldRateLimit(context.Background(), &req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
