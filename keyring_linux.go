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

// TODO: Add bindings for KDE Wallet.

// #cgo pkg-config: glib-2.0 gnome-keyring-1
// #include <stdlib.h>
// #include "include/gnome-keyring/keyring.c"
import "C"

import (
	"unsafe"
)

// A client that talks to Gnome's Keyring service.
type GnomeKeyring struct{}

// Returns true if gnome keyring service is available.
func (k *GnomeKeyring) IsAvailable() bool {
	return C.gnome_keyring_isavailable() > 0
}

// Gets the password for the specified service and username.
func (k *GnomeKeyring) Get(service, username string) (string, error) {
	s := C.CString(service)
	u := C.CString(username)
	p := C.CString("")
	defer func() {
		C.free(unsafe.Pointer(s))
		C.free(unsafe.Pointer(u))
	}()
	if code := C.gnome_keyring_get(s, u, &p); code != 0 {
		C.free(unsafe.Pointer(p))
		return "", k.convertResultToErr(code)
	}
	return C.GoString(p), nil
}

// Sets a password for the specified service and username.
func (k *GnomeKeyring) Set(service, username, password string) error {
	s := C.CString(service)
	u := C.CString(username)
	p := C.CString(password)

	defer func() {
		C.free(unsafe.Pointer(s))
		C.free(unsafe.Pointer(u))
		C.free(unsafe.Pointer(p))
	}()
	return k.convertResultToErr(C.gnome_keyring_set(s, u, p))
}

// Deletes the password belongs to the specified service and username
// if there exists one. Silently returns if there is no such item.
func (k *GnomeKeyring) Delete(service, username string) error {
	s := C.CString(service)
	u := C.CString(username)

	defer func() {
		C.free(unsafe.Pointer(s))
		C.free(unsafe.Pointer(u))
	}()
	return k.convertResultToErr(C.gnome_keyring_del(s, u))
}

// Converts C.GnomeKeyringResult to an error
func (k *GnomeKeyring) convertResultToErr(result C.GnomeKeyringResult) error {
	switch result {
	case 0:
		return nil
	case C.GNOME_KEYRING_RESULT_NO_MATCH:
		return ErrNotFound
	default:
		return ErrUnknown
	}
}

func init() {
	Register("gnome-keyring", &GnomeKeyring{})
}
