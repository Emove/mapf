package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kratos_grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func NewClientConnection(endpoint string, timeout time.Duration) (*grpc.ClientConn, error) {
	return kratos_grpc.Dial(
		context.Background(),
		kratos_grpc.WithEndpoint(endpoint),
		kratos_grpc.WithTimeout(timeout),
		kratos_grpc.WithOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		kratos_grpc.WithMiddleware(
			recovery.Recovery(),
		))
}
