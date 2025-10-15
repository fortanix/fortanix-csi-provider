package server

import (
	"context"
	"fmt"
	"log"

	"github.com/fortanix/fortanix-csi-provider/internal/config"
	provider "github.com/fortanix/fortanix-csi-provider/internal/provider"
	pb "github.com/fortanix/fortanix-csi-provider/internal/v1alpha1"
	"github.com/fortanix/fortanix-csi-provider/internal/version"
)

var _ pb.CSIDriverProviderServer = (*Server)(nil)

type Server struct {
	DsmApiKey   string
	DsmEndpoint string
}

func (s *Server) Version(context.Context, *pb.VersionRequest) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{
		Version:        "v1alpha1",
		RuntimeName:    "fortanix-csi-provider",
		RuntimeVersion: version.BuildVersion,
	}, nil
}

func (s *Server) Mount(ctx context.Context, req *pb.MountRequest) (*pb.MountResponse, error) {
	cfg, err := config.Parse(
		req.Attributes,
		req.TargetPath,
		req.Permission,
	)
	if err != nil {
		log.Printf("Error parsing config: %v", err)
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Parameters.DsmEndpoint == "" || cfg.Parameters.DsmApiKey == "" {
		log.Println("SecretProviderClass not found or invalid")
		return nil, fmt.Errorf("SecretProviderClass not found or invalid")
	}

	provider := provider.NewProvider()
	resp, err := provider.HandleMountRequest(ctx, cfg)
	if err != nil {
		log.Printf("Error handling mount request: %v", err)
		return nil, fmt.Errorf("error making mount request: %w", err)
	}

	return resp, nil
}
