/* Copyright (c) Fortanix, Inc.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"log/slog"
	"net/http"

	"github.com/fortanix/sdkms-client-go/sdkms"
	"github.com/pkg/errors"

	"github.com/fortanix/fortanix-csi-provider/internal/config"
)

type SecretClient struct {
	*sdkms.Client
}

func NewSecretClient(parameters config.SpcParameters) (*SecretClient, error) {
	if parameters.DsmEndpoint == "" {
		slog.Error("Endpoint empty")
		return nil, errors.Errorf("Could not find an endpoint")
	}
	if parameters.ApiKey == "" {
		slog.Error("Api Key empty")
		return nil, errors.Errorf("Could not find an api key")
	}
	client := sdkms.Client{
		HTTPClient: http.DefaultClient,
		Auth:       sdkms.APIKey(parameters.ApiKey),
		Endpoint:   parameters.DsmEndpoint,
	}
	return &SecretClient{&client}, nil
}
