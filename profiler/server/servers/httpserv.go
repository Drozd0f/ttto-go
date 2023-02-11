package servers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	profilerv1 "github.com/Drozd0f/ttto-go/gen/proto/go/profiler/v1"
	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewHttpServer(ctx context.Context, logg *zap.Logger, conf *config.Config) error {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))

	dos := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := profilerv1.RegisterProfilerServiceHandlerFromEndpoint(ctx, mux, conf.GrpcAdress(), dos); err != nil {
		return fmt.Errorf("register profiler service handler: %w", err)
	}

	s := http.Server{
		Handler:      mux,
		Addr:         conf.HttpAdress(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	logg.Info("started serving profiler http service")

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil && !errors.Is(err, context.DeadlineExceeded) {
			log.Print("profiler server shutdown:", err)
		}

		log.Print("stoped serving profiler http service")
	}()

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("http serve: %w", err)
	}

	return nil
}
