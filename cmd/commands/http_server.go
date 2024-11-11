package commands

import (
	"fin_data_processing/internal/config"
	"fin_data_processing/internal/handler/httphandler"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunHttp(cfg *config.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/api/ping", ping).Methods("GET")
	r.HandleFunc("/api/users", httphandler.AddUser).Methods("POST")
	r.HandleFunc("/api/users", httphandler.UsersList).Methods("GET")
	r.HandleFunc("/api/users/{userId}", httphandler.UserOne).Methods("GET")
	r.HandleFunc("/api/users/{id}", httphandler.UserDelete).Methods("DELETE")

	r.HandleFunc("/api/users/security-fulfils", httphandler.AddSecurityFulfil).Methods("POST")
	r.HandleFunc("/api/users/security-fulfils", httphandler.SecurityFulfilsList).Methods("GET")
	r.HandleFunc("/api/users/security-fulfils/{id}", httphandler.SecurityFulfilOne).Methods("GET")
	r.HandleFunc("/api/users/security-fulfils/{id}", httphandler.SecurityFulfilDelete).Methods("DELETE")

	go func() {
		err := http.ListenAndServe(cfg.ServerAddress, r)
		if err != nil {
			slog.Info("Error starting server:", err)
			// Ждем несколько секунд перед перезапуском
			time.Sleep(5 * time.Second)
			slog.Info("Error starting server:", err)
		}

		slog.Info("Error starting server:", err)
	}()

	slog.Info("Сервер http запущен")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	slog.Info("Shutting down server...")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok!")
}
