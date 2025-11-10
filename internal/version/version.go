/* Copyright (c) Fortanix, Inc.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package version

import (
	"encoding/json"
)

const minDriverVersion = "v0.0.21"

var (
	BuildDate    string
	BuildVersion string
	GoVersion    string
)

// providerVersion holds current provider version
type providerVersion struct {
	Version          string `json:"version"`          // Version of the binary.
	BuildDate        string `json:"buildDate"`        // The date the binary was built.
	GoVersion        string `json:"goVersion"`        // Version of Go the binary was built with.
	MinDriverVersion string `json:"minDriverVersion"` // Minimum driver version the provider works with.
}

func GetVersion() (string, error) {
	pv := providerVersion{
		Version:          BuildVersion,
		BuildDate:        BuildDate,
		GoVersion:        GoVersion,
		MinDriverVersion: minDriverVersion,
	}

	res, err := json.Marshal(pv)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
