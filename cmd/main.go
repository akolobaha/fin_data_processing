package main

import (
	"context"
	pb "fin_data_processing/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedDataManagementServiceServer
}

func (s *server) GetQuotes(ctx context.Context, TickerReq *pb.TickerRequest) (*pb.TickerResponse, error) {
	return &pb.TickerResponse{Price: TickerReq.Price, Name: TickerReq.Name, Time: TickerReq.Time}, nil
	// TODO: поймать данные и записать в БД
	// TODO: В миграцию положить данные из https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities.json
	//return nil, status.Errorf(codes.Unimplemented, "method GetQuotes has been implemented")
}

//func (s *server) SendMessage(ctx context.Context, msg *pb.TickerRequest) (*pb.TickerResponse, error) {
//	log.Printf("Received message: %s", msg.Name)
//	return (*pb.TickerResponse)(msg), nil
//}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataManagementServiceServer(s, &server{})
	log.Println("Server is running on port 50052")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
