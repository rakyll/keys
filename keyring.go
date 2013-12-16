// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package keyring

import (
	"errors"
)

type Keyring interface {
	// Returns if keyring backend is available on the
	// current platform where program is running.
	IsAvailable() bool

	// Gets the password for the service and username if exists.
	Get(service, username string) (string, error)

	// Sets a password for the service and username.
	Set(service, username, password string) error

	// Deletes a password belongs to the service and username.
	Delete(service, username string) error
}

var (
	// A list of all registered Keyring backends.
	registered map[string]Keyring
)

// Returns the first keyring backend available on the platform.
// If no keyring is available, returns an error.
func New() (Keyring, error) {
	for _, keyring := range registered {
		if keyring.IsAvailable() {
			return keyring, nil
		}
	}
	return nil, errors.New("no available keyring found")
}

// Registers a Keyring backend.
func Register(key string, keyring Keyring) {
	if registered == nil {
		registered = make(map[string]Keyring)
	}
	registered[key] = keyring
}
