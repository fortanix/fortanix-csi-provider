/* Copyright (c) Fortanix, Inc.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type FortanixConfig struct {
	endpoint string
	apikey   string
}

type Provider struct {
	cache    map[string]string
	versions map[string]string
}
type SpcParameters struct {
	DsmEndpoint string
	ApiKey      string
	Secrets     []Secret
}

func NewFortanixConfig() *FortanixConfig {
	return &FortanixConfig{}
}

type Parameters struct {
	DsmApiKey           string `json:"dsmApikey"`
	DsmEndpoint         string `json:"dsmEndpoint"`
	Secrets             []Secret
	PodName             string
	ServiceAccountName  string `json:"csi.storage.k8s.io/serviceAccount.name"`
	Namespace           string `json:"csi.storage.k8s.io/pod.namespace"`
	ServiceAccountToken string
	UID                 string `json:"csi.storage.k8s.io/pod.uid"`
}
type Config struct {
	Parameters
	TargetPath     string
	FilePermission os.FileMode
}

type Secret struct {
	SecretName     string
	FilePermission os.FileMode
}

type FlagsConfig struct {
	Endpoint    string
	DsmEndpoint string
	DsmApiKey   string
	Version     bool
	HealthAddr  string
}

func Parse(
	parametersStr, targetPath, permissionStr string,
) (Config, error) {
	config := Config{
		TargetPath: targetPath,
	}
	var err error
	config.Parameters, err = parseParameters(parametersStr)
	if err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal([]byte(permissionStr), &config.FilePermission); err != nil {
		return Config{}, err
	}

	if err := config.validate(); err != nil {
		return Config{}, err
	}

	return config, nil
}

func parseParameters(parametersStr string) (Parameters, error) {
	var params map[string]string
	if err := json.Unmarshal([]byte(parametersStr), &params); err != nil {
		log.Println("Failed to Unmarshal Mount request- Get attributes")
		return Parameters{}, fmt.Errorf("Failed to Unmarshal Mount request - Get attributes")
	}

	var parameters Parameters
	parameters.DsmApiKey = os.Getenv("FORTANIX_API_KEY")
	parameters.DsmEndpoint = params["dsmEndpoint"]
	parameters.PodName = params["csi.storage.k8s.io/pod.name"]
	parameters.UID = params["csi.storage.k8s.io/pod.uid"]
	parameters.Namespace = params["csi.storage.k8s.io/pod.namespace"]
	parameters.ServiceAccountName = params["csi.storage.k8s.io/serviceAccount.name"]
	if parameters.DsmEndpoint == "" {
		parameters.DsmEndpoint = os.Getenv("FORTANIX_DSM_ENDPOINT")
	}
	var secrets []Secret
	secretStr := params["objects"]
	secretEntries := strings.Split(secretStr, "- secretName: ")
	for _, entry := range secretEntries {
		secretName := strings.Trim(strings.TrimSpace(entry), "\"")
		if secretName != "" {
			secrets = append(secrets, Secret{SecretName: secretName})
		}
	}
	parameters.Secrets = secrets
	return parameters, nil
}

func (c *Config) validate() error {
	// Some basic validation checks.
	if c.TargetPath == "" {
		return errors.New("missing target path field")
	}
	if c.Parameters.DsmApiKey == "" {
		return errors.New("missing API key - must be set via FORTANIX_API_KEY environment variable")
	}
	if c.Parameters.DsmEndpoint == "" {
		return errors.New("missing DSM endpoint")
	}
	if len(c.Parameters.Secrets) == 0 {
		return errors.New("no secrets configured - the provider will not read any secret material")
	}

	objectNames := map[string]struct{}{}
	conflicts := []string{}
	for _, secret := range c.Parameters.Secrets {
		if _, exists := objectNames[secret.SecretName]; exists {
			conflicts = append(conflicts, secret.SecretName)
		}

		objectNames[secret.SecretName] = struct{}{}
	}

	if len(conflicts) > 0 {
		return fmt.Errorf("each `secretName` within a SecretProviderClass must be unique, "+
			"but the following keys were duplicated: %s", strings.Join(conflicts, ", "))
	}

	return nil
}
