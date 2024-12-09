package entities

import (
	"context"
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/monitoring"
	pb "fin_data_processing/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"time"
)

type Target struct {
	ID                 int64
	Ticker             string
	ValuationRatio     string
	Value              float64
	FinancialReport    string
	Achieved           bool
	NotificationMethod string
}

type TargetUser struct {
	Target      Target
	User        User
	ResultValue float64
}

func FetchTargets(ticker string, cfg *config.Config) []TargetUser {
	conn, err := grpc.NewClient(cfg.GetGrpc(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTargetsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TargetRequest{Ticker: ticker}

	response, err := client.GetTargets(ctx, req)
	if err != nil {
		monitoring.ProcessingErrorCount.WithLabelValues("Could not fetch targets").Inc()
		slog.Error("could not get targets: %v", "error", err)
		return []TargetUser{}
	}

	targetsUsers := make([]TargetUser, len(response.Targets))

	for i, t := range response.Targets {
		targetsUsers[i].Target = Target{
			ID:                 t.Id,
			Ticker:             t.Ticker,
			ValuationRatio:     t.ValuationRatio,
			Value:              float64(t.Value),
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

func SetTargetAchieved(targetID int64, achieved bool) TargetUser {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTargetsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TargetAchievedRequest{Id: targetID, Achieved: achieved}

	response, err := client.SetTargetAchieved(ctx, req)
	if err != nil {
		slog.Error("could not set target achieved: %v", "error", err)
	}

	return TargetUser{
		Target: Target{
			ID:                 response.Id,
			Ticker:             response.Ticker,
			ValuationRatio:     response.ValuationRatio,
			FinancialReport:    response.FinancialReport,
			Achieved:           response.Achieved,
			NotificationMethod: response.NotificationMethod,
		},
		User: User{
			ID:       response.User.Id,
			Name:     response.User.Name,
			Email:    response.User.Email,
			Telegram: response.User.Telegram,
		},
	}
}
