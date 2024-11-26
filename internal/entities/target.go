package entities

import (
	"context"
	pb "fin_data_processing/pkg/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Target struct {
	Id                 int64
	Ticker             string
	ValuationRatio     string
	Value              float64
	FinancialReport    string
	Achieved           bool
	NotificationMethod string
}

type TargetUser struct {
	Target Target
	User   User
}

func FetchTargets() []TargetUser {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure()) // Убедитесь, что порт совпадает с вашим сервером

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTargetsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TargetRequest{}

	response, err := client.GetTargets(ctx, req)
	if err != nil {
		log.Fatalf("could not get targets: %v", err)
	}

	targetsUsers := make([]TargetUser, len(response.Targets))

	for i, t := range response.Targets {
		targetsUsers[i].Target = Target{
			Id:                 t.Id,
			Ticker:             t.Ticker,
			ValuationRatio:     t.ValuationRatio,
			FinancialReport:    t.FinancialReport,
			Achieved:           t.Achieved,
			NotificationMethod: t.NotificationMethod,
		}
		targetsUsers[i].User = User{
			ID:       t.User.Id,
			Name:     t.User.Name,
			Email:    t.User.Email,
			Telegram: t.User.Telegram,
		}
	}

	return targetsUsers
}