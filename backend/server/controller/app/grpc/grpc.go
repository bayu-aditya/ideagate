package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/bayu-aditya/ideagate/backend/model/gen-go/core/application"
	"github.com/bayu-aditya/ideagate/backend/model/gen-go/dashboard"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Action(c *cli.Context) error {
	lisGrpc, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	lisGrpcWeb, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register gRPC services
	dashboard.RegisterDashboardServiceServer(s, &DashboardServiceServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("gRPC server is running on port 50051")
		if err := s.Serve(lisGrpc); err != nil {
			fmt.Errorf("failed to serve: %v", err)
		}
	}()

	wrappedGrpcWeb := grpcweb.WrapServer(s)
	httpServer := &http.Server{
		Handler: middleware(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if wrappedGrpcWeb.IsGrpcWebRequest(req) {
				wrappedGrpcWeb.ServeHTTP(resp, req)
				return
			}
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(resp, req)
		})),
	}

	go func() {
		fmt.Println("gRPC web server is running on port 50052")
		if err := httpServer.Serve(lisGrpcWeb); err != nil {
			fmt.Errorf("failed to serve: %v", err)
		}
	}()

	select {}

	return nil
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Disposition, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-grpc-web")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Max-Age", "600")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type DashboardServiceServer struct {
	dashboard.UnimplementedDashboardServiceServer
}

func (s *DashboardServiceServer) GetListApplication(context.Context, *dashboard.GetListApplicationRequest) (*dashboard.GetListApplicationResponse, error) {
	return &dashboard.GetListApplicationResponse{
		Applications: []*application.Application{
			{Id: "1", Name: "App 1"},
			{Id: "2", Name: "App 2"},
		},
	}, nil
}
