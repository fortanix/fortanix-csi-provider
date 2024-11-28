package provider

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fortanix/sdkms-client-go/sdkms"

	"github.com/kiran-m-kannur/fortanix-csi-provider/internal/client"
	"github.com/kiran-m-kannur/fortanix-csi-provider/internal/config"
	pb "github.com/kiran-m-kannur/fortanix-csi-provider/internal/v1alpha1"
)

type provider struct{}

func NewProvider() *provider {
	p := &provider{}
	return p
}

func (p *provider) getSecret(
	ctx context.Context,
	client *client.SecretClient,
	secretConfig config.Secret,
) ([]byte, error) {
	secretName := secretConfig.SecretName
	sobjectreq := sdkms.SobjectByName(secretName)
	sobject, err := client.ExportSobject(ctx, *sobjectreq)
	if err != nil {
		log.Printf("Error! Could not fetch the Sobject %v: %v", secretName, err)
		return nil, err
	}
	if sobject.Value == nil {
		return nil, fmt.Errorf("Sobject %v has no value", secretName)
	}
	return *sobject.Value, nil
}

func (p *provider) HandleMountRequest(
	ctx context.Context,
	cfg config.Config,
) (*pb.MountResponse, error) {
	authconfig := config.SpcParameters{
		DsmEndpoint: cfg.Parameters.DsmEndpoint,
		ApiKey:      cfg.Parameters.DsmApiKey,
	}
	client, err := client.NewSecretClient(authconfig)
	if err != nil {
		log.Fatalf("Error creating a new Client :%v", err)
		return nil, err
	}

	var files []*pb.File
	var objectVersions []*pb.ObjectVersion

	for _, secret := range cfg.Parameters.Secrets {
		log.Println("Fetching :", secret.SecretName)

		content, err := p.getSecret(ctx, client, secret)
		if err != nil {
			return nil, err
		}

		hash := sha256.Sum256(content)
		objectVersion := &pb.ObjectVersion{
			Id: hex.EncodeToString(hash[:]),
		}
		filePermission := int32(cfg.FilePermission)
		if secret.FilePermission != 0 {
			filePermission = int32(secret.FilePermission)
		}
		files = append(
			files,
			&pb.File{Path: secret.SecretName, Mode: filePermission, Contents: content},
		)
		objectVersions = append(objectVersions, objectVersion)

		log.Println(
			"secret added to mount response",
			"directory",
			cfg.TargetPath,
			"file:",
			secret.SecretName,
		)
	}
	return &pb.MountResponse{
		Files:         files,
		ObjectVersion: objectVersions,
	}, nil
}

func generateObjectVersion(
	secret config.Secret,
	hmacKey []byte,
	content []byte,
) (*pb.ObjectVersion, error) {
	if hmacKey == nil {
		return &pb.ObjectVersion{
			Id:      secret.SecretName,
			Version: "",
		}, nil
	}
	hash := hmac.New(sha256.New, hmacKey)
	cfg, err := json.Marshal(secret)
	if err != nil {
		return nil, err
	}
	if _, err := hash.Write(cfg); err != nil {
		return nil, err
	}
	if _, err := hash.Write(content); err != nil {
		return nil, err
	}
	return &pb.ObjectVersion{
		Id:      secret.SecretName,
		Version: base64.URLEncoding.EncodeToString(hash.Sum(nil)),
	}, nil
}
