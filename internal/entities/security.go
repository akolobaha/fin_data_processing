package entities

import (
	"context"
	pb "fin_data_processing/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Securities []Security

type Security struct {
	Ticker    string `gorm:"ticker"`
	Shortname string `gorm:"shortname"`
	Secname   string `gorm:"secname"`
}

func FetchSecurities() Securities {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure()) // Убедитесь, что порт совпадает с вашим сервером
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTickersServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TickersRequest{}

	response, err := client.GetMultipleTickers(ctx, req)
	if err != nil {
		log.Fatalf("could not get tickers: %v", err)
	}

	securities := make(Securities, len(response.Tickers))

	for i, ticker := range response.Tickers {
		securities[i].Shortname = ticker.Shortname
		securities[i].Secname = ticker.Name
		securities[i].Ticker = ticker.Ticker
	}
	return securities
}
