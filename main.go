package main

import (
	"os"

	"dummy_ratelimit/grpc"
	"dummy_ratelimit/ratelimit"

	log "github.com/sirupsen/logrus"
)

var (
	listenAddr = ":8080"
)

func init() {
	setupLogger()
	setupListenAddress()
}

func main() {
	grpcServer := grpc.NewGRPCServer(listenAddr)
	ratelimitServer := ratelimit.NewRateLimitServer()
	grpcServer.Register(ratelimitServer)
	grpcServer.Run()
}

func setupLogger() {
	if path := os.Getenv("LOG_TO_FILE"); path != "" {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	}
	log.SetFormatter(new(log.JSONFormatter))
}

func setupListenAddress() {
	if val := os.Getenv("LISTEN_ADDRESS"); val != "" {
		listenAddr = val
	}
}
