package seaotter

import (
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

type seaotterGRPCImpl struct {
	host    string
	authKey string
	conn    *grpc.ClientConn
}

// NewSeaotterServiceGRPC constructor
func NewSeaotterServiceGRPC(host string, authKey string) Seaotter {

	if u, _ := url.Parse(host); u.Host != "" {
		host = u.Host
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &seaotterGRPCImpl{
		host:    host,
		authKey: authKey,
		conn:    conn,
	}
}
