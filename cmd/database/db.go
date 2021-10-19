package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"publisher/internal/database"
	dbsvc "publisher/pkg/database"
	"publisher/pkg/database/endpoints"
	"publisher/pkg/database/transport"
	"syscall"

	pb "publisher/api/v1/pb/db"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/oklog/run"
	"google.golang.org/grpc"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

var (
	logger   log.Logger
	httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
	grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
)

func main() {
	var (
		service     = dbsvc.NewService()
		endpointSet = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(endpointSet)
		grpcServer  = transport.NewGRPCServer(endpointSet)
	)

	var g run.Group
	{
		// HTTP Listener by Go Kit Http Handler
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// GRPC Listener by Go Kit gRPC Server
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterDatabaseServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// try connect postgreSQL database
	db, err := database.Init(database.DefaultDatabase, database.DefaultHost, database.DefaultPort, database.DefaultDBUser, database.DefaultPassword, database.DefaultTimeZone)
	defer func() {
		sqlDB, err := db.DB()
		err = sqlDB.Close()
		if err != nil {
			logger.Log("ERROR::Failed to close the database connection", err.Error())
		}
	}()
	if err != nil {
		logger.Log(fmt.Sprintf("FATAL: failed to load db with error: %s", err.Error()))
	}
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
