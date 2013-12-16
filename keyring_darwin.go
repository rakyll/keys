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
	"fmt"
	"os/exec"
)

const (
	execPathKeychain = "/usr/bin/security"
)

type Keychain struct{}

func (*Keychain) IsAvailable() bool {
	return exec.Command(execPathKeychain).Run() != exec.ErrNotFound
}

func (k *Keychain) Get(service, username string) (string, error) {
	out, err := exec.Command(
		execPathKeychain,
		"find-generic-password",
		"-g",
		"-s", service,
		"-a", username).Output()
	return k.handleGetOut(out, err)
}

func (k *Keychain) Set(service, username, password string) error {
	existing, _ := k.Get(service, username)
	if existing != "" {
		if err := k.Delete(service, username); err != nil {
			return errors.New("password exists for service, username, can't delete")
		}
	}
	return exec.Command(
		execPathKeychain,
		"add-generic-password",
		"-s", service,
		"-a", username,
		"-w", password).Run()
}

func (k *Keychain) Delete(service, username string) error {
	return exec.Command(
		execPathKeychain,
		"delete-generic-password",
		"-s", service,
		"-a", username).Run()
}

func (k *Keychain) handleGetOut(out []byte, err error) (string, error) {
	if err != nil {
		// handle cases where no record exists
		return "", err
	}
	return fmt.Sprintf("%s", out), nil
}

func init() {
	Register("macosx-keyring", &Keychain{})
}
