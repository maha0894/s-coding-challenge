package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/maha0894/s-coding-challenge/pkg/application"
	"github.com/maha0894/s-coding-challenge/pkg/repository"
	"github.com/maha0894/s-coding-challenge/pkg/transport"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts HTTP server",
	Long:  "Starts HTTP server to handle requests",
	Run:   func(cmd *cobra.Command, args []string) { runServer() },
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

func runServer() {
	done := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	router := mux.NewRouter()
	repo, err := repository.New()
	if err != nil {
		log.Fatalf("Failed to initialise repository: %v", err)
	}
	service := application.NewService(repo)
	handler := transport.NewUserHandler(service)
	handler.RegisterRoutes(router)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: router,
	}

	go func() {
		<-signals
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Fatalf("could not shutdown http server gracefully: %v", err)
		}
		close(done)
	}()

	go func() {
		log.Print("Starting http server...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run http server: %v", err)
		}
	}()

	<-done
}
