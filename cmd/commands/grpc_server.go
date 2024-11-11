package commands

import (
	"context"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	pb "fin_data_processing/pkg/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

type server struct {
	pb.UnimplementedDataManagementServiceServer
}

func (s *server) GetMultipleQuotes(ctx context.Context, TickerReq *pb.MultipleTickerRequest) (*pb.TickerResponse, error) {
	err := entities.InsertQuotes(TickerReq)
	if err != nil {
		return nil, err
	}

	return &pb.TickerResponse{Response: nil}, nil
}

func RunGrpc(cfg *config.Config) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataManagementServiceServer(s, &server{})

	go func() {
		slog.Info(fmt.Sprintf("Server is running on port: %s", cfg.GrpcPort))

		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve: ", err.Error())
		}
	}()

}
