package log

import (
	"fin_data_processing/internal/monitoring"
	"fmt"
	"log/slog"
)

func Error(additionalMessage string, err error) {
	if err != nil {
		msg := fmt.Sprintf("%s: %s", additionalMessage, err.Error())
		monitoring.ProcessingErrorCount.WithLabelValues(msg).Inc()
		slog.Error(msg)
	}
}

func Info(message string) {
	monitoring.ProcessingSuccessCount.WithLabelValues(message).Inc()
	slog.Info(message)
}
