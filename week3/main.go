package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"

	db "week3/data"
	"week3/service"
	gw "week3/transport"
)

// Config is configuration for Server
type Config struct {
	GRPCPort   string
	HTTPPort   string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc", "4000", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http", "3000", "HTTP port to bind")
	flag.StringVar(&cfg.DBHost, "host", "", "Database host")
	flag.StringVar(&cfg.DBUser, "user", "", "Database user")
	flag.StringVar(&cfg.DBPassword, "password", "", "Database password")
	flag.StringVar(&cfg.DBName, "db", "", "Database name")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	api := service.NewFeedbackService(db.Repository)

	go func() {
		gw.RunGrpcServer(ctx, api, cfg.GRPCPort)
	}()

	return gw.RunHTTPServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
