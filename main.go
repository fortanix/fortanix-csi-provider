package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	providerserver "github.com/kiran-m-kannur/fortanix-csi-provider/internal/server"
	pb "github.com/kiran-m-kannur/fortanix-csi-provider/internal/v1alpha1"
	"github.com/kiran-m-kannur/fortanix-csi-provider/internal/version"
)

func realMain() error {
	var (
		endpoint = flag.String(
			"endpoint",
			"/provider/fortanix-csi-provider.sock",
			"path to socket on which to listen for driver gRPC calls",
		)
		selfVersion = flag.Bool("version", false, "prints the version information")
		dsmAddr     = flag.String("dsm-address", "https://api.smartkey.io", "Fortanix API URL")
		healthAddr  = flag.String(
			"health-address",
			":8080",
			"configure http listener for reporting health",
		)
	)

	flag.Parse()

	if *selfVersion {
		v, err := version.GetVersion()
		if err != nil {
			return fmt.Errorf("failed to print version, err: %w", err)
		}
		_, err = fmt.Println(v)
		return err
	}

	log.Print("Creating new gRPC server")
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				startTime := time.Now()
				log.Printf("Processing unary gRPC call grpc.method: %v", info.FullMethod)
				resp, err := handler(ctx, req)
				log.Printf(
					"Finished unary gRPC call grpc.method: %v, grpc.time: %v, grpc.code: %v",
					info.FullMethod,
					time.Since(startTime),
					status.Code(err),
				)
				if err != nil {
					log.Printf("Error: %v", err.Error())
				}
				log.Print("Finished unary gRPC call")
				return resp, err
			},
		),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-c
		log.Printf("Caught signal %s, shutting down", sig)
		server.GracefulStop()
	}()

	listener, err := listen(*endpoint)
	if err != nil {
		return err
	}
	defer listener.Close()

	s := &providerserver.Server{
		DsmEndpoint: *dsmAddr,
	}
	pb.RegisterCSIDriverProviderServer(server, s)

	// Create health handler
	mux := http.NewServeMux()
	ms := http.Server{
		Addr:    *healthAddr,
		Handler: mux,
	}
	defer func() {
		err := ms.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("Error shutting down health handler, err: %v", err.Error())
		}
	}()

	mux.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Start health handler
	go func() {
		log.Printf("Starting health handler, addr: %v", *healthAddr)
		if err := ms.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error with health handler, error: %v", err.Error())
		}
	}()

	log.Print("Starting gRPC server")
	err = server.Serve(listener)
	if err != nil {
		return fmt.Errorf("error running gRPC server: %v", err.Error())
	}

	return nil
}

func listen(endpoint string) (net.Listener, error) {
	_, err := os.Stat(endpoint)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to check for existence of unix socket: %v", err.Error())
	} else if err == nil {
		log.Printf("Cleaning up pre-existing file at unix socket location, endpoint: %v", endpoint)
		err = os.Remove(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to clean up pre-existing file at unix socket location: %v", err.Error())
		}
	}

	log.Printf("Opening unix socket, endpoint %v", endpoint)
	listener, err := net.Listen("unix", endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on unix socket at %s: %v", endpoint, err.Error())
	}

	return listener, nil
}

func main() {
	err := realMain()
	if err != nil {
		log.Fatalf("Error running provider: %v", err.Error())
		os.Exit(1)
	}
}
