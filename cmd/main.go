package main

import (
	"context"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/entities"
	pb "fin_data_processing/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

const defaultEnvFilePath = ".env"

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

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)
	config.InitDbConnectionString(cfg)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataManagementServiceServer(s, &server{})
	slog.Info("Server is running on port 50052")

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve: ", err.Error())
	}
}
