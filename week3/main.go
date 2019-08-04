package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/golang/glog"

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

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	// param := "parseTime=true"

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
	// 	cfg.DatastoreDBUser,
	// 	cfg.DatastoreDBPassword,
	// 	cfg.DatastoreDBHost,
	// 	cfg.DatastoreDBSchema,
	// 	param)
	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	return fmt.Errorf("failed to open database: %v", err)
	// }
	// defer db.Close()

	api := service.NewFeedbackService()

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
