package grpc

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/muharik19/boiler-plate-grpc/api/grpc/api"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/database"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/elasticsearch"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	"golang.org/x/sync/errgroup"
)

func Init() {
	ctx := context.Background()

	// Connect to the database postgres
	_, err := database.InitBun(ctx)
	if err != nil {
		logger.Errf("db connection %s", err.Error())
	}

	elasticsearch.Init()

	g, _ := errgroup.WithContext(ctx)
	var servers []*http.Server
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			log.Printf("received signal: %s\n", sig)
			for i, s := range servers {
				if err := s.Shutdown(ctx); err != nil {
					log.Printf("error shutting down server %d: %v", i, err)
					panic(err)
				}
			}
			os.Exit(1)
		}
		return nil
	})

	g.Go(func() error { return api.NewGrpcServer() })
	g.Go(func() error { return api.NewHttpServer() })
	err = g.Wait()
	if err != nil {
		panic(err)
	}
	return
}
