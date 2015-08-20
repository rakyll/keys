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

// Only exists because cgo doesn't support variable length argument lists,
// which is required for libgnome-keyring.

#ifndef GNOME_KEYRING_WRAPPER_C
#define GNOME_KEYRING_WRAPPER_C

#include "gnome-keyring.h"

// TODO: move this under /src or somewhere more appropriate
const GnomeKeyringPasswordSchema GNOME_KEYRING_DEFAULT_SCHEMA = {
  GNOME_KEYRING_ITEM_GENERIC_SECRET,
  {
    { "service",  GNOME_KEYRING_ATTRIBUTE_TYPE_STRING },
    { "username", GNOME_KEYRING_ATTRIBUTE_TYPE_STRING },
    { NULL, 0 }
  }
};

int
gnome_keyring_isavailable() {
  // gnome_key
  return gnome_keyring_is_available();
}

GnomeKeyringResult
gnome_keyring_get(const char *service,
                  const char *username,
                  char **password_ptr) {
  return gnome_keyring_find_password_sync(
      &GNOME_KEYRING_DEFAULT_SCHEMA,
      password_ptr,
      "service",  service,
      "username", username,
      NULL);
}

GnomeKeyringResult
gnome_keyring_set(const char *service,
                  const char *description,
                  const char *username,
                  const char *password) {
  return gnome_keyring_store_password_sync(
      &GNOME_KEYRING_DEFAULT_SCHEMA,
      NULL,
      description,
      password,
      "service",  service,
      "username", username,
      NULL);
}

GnomeKeyringResult
gnome_keyring_del(const char *service,
                  const char *username) {
  return gnome_keyring_delete_password_sync(
      &GNOME_KEYRING_DEFAULT_SCHEMA,
      "service",  service,
      "username", username,
      NULL);
}

#endif /* GNOME_KEYRING_WRAPPER_C */
