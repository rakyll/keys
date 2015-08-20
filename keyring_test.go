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

import "testing"

// declare testing data
var user string = "username"
var password string = "testpassword"
var service string = "keyring-unit-test"

func TestNew(t *testing.T) {
	// Any Error declaring a new keyring should result
	// in an immediate failure as all other tests will fail.
	_, err := New()
	if err != nil {
		t.Fatal("New() Error (Fatal): ", err)
	}
}

func TestIsAvailable(t *testing.T) {
	// As TestNew() is run first and fatals on an error,
	// this test should indicate a failure of the function
	kr, _ := New()
	i := kr.IsAvailable()
	if i != true {
		t.Error("IsAvailable() Error: Should return true, received: ", i)
	}
}

func TestSet(t *testing.T) {
	kr, _ := New()
	err := kr.Set(service, user, password)
	if err != nil {
		t.Error("Set() Error: ", err)
	}
	if _, err := kr.Get(service, user); err == ErrNotFound {
		t.Error("Set() Error on Get(): ", err)
	}
}

func TestGet(t *testing.T) {
	kr, _ := New()
	p, err := kr.Get(service, user)
	if err != nil {
		t.Error("Get() Error: ", err)
		// Not Matching may be acceptable if this is the initial Get.
	} else if p != password {
		t.Errorf("Get() Error: Password does not match the test data. Expected %s but got %s", password, p)
	}
}

func TestDelete(t *testing.T) {
	kr, _ := New()
	err := kr.Delete(service, user)
	if err != nil {
		t.Error("Delete() Error: ", err)
	}
	_, err = kr.Get(service, user)
	if err != ErrNotFound {
		t.Error("Delete() Error: entry for %s:%s not successfully removed.")
	}
}
