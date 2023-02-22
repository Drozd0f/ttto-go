package servers

import (
	"context"
	"fmt"
	"net"
	"time"

	profilerv1 "github.com/Drozd0f/ttto-go/gen/proto/go/profiler/v1"
	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/Drozd0f/ttto-go/profiler/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewGrpcServer(ctx context.Context, logg *zap.Logger, conf *config.Config, serv *server.Server) error {
	grpcServer := grpc.NewServer()

	profilerv1.RegisterProfilerServiceServer(grpcServer, serv)

	logg.Info("started serving profiler grpc service")

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		go func() {
			<-ctx.Done()

			grpcServer.Stop()
			logg.Info("stoped profiler grpc server")
		}()

		grpcServer.GracefulStop()
	}()

	lis, err := net.Listen("tcp", conf.GrpcAdress())
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("gprc serve: %w", err)
	}

	return nil
}
