package entities

import (
	"fin_data_processing/db"
	pb "fin_data_processing/pkg/grpc"
	"log/slog"
)

func InsertQuotes(TickerReq *pb.MultipleTickerRequest) error {
	db := db.GetDbConnection()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	// Prepare statement for batch insert
	stmt, err := tx.Prepare("INSERT INTO quotes(ticker, price, time, seq_num) VALUES ($1, $2, $3, $4)")
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer stmt.Close()

	for _, ticker := range TickerReq.Tickers {

		_, err := stmt.Exec(ticker.Name, ticker.Price, ticker.Time.AsTime(), ticker.SeqNum)
		if err != nil {
			tx.Rollback() // Rollback if there is an error
			slog.Error(err.Error())
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
