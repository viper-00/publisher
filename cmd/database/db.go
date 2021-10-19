package main

import (
	"fmt"
	"go-kit-example/pkg/database"
	"go-kit-example/pkg/database/endpoints"
	"go-kit-example/pkg/database/transport"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/oklog/run"
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
		service     = database.NewService()
		endpointSet = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(endpointSet)
	)

	var g run.Group
	{
		// Http Listener by Go Kit Http Handler
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
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
