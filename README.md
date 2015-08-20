# Keys

Keys is a cross-platform keyring library for golang, utilizing gnome-kerying onLinux or the Mac OSX Keychain to Get, Set, or Delete passwords.

## Dependencies
- libglib2.0-dev (gnome-keyring support)

## Example Usage
```golang
package main

import (
"github.com/rakyll/keys"
"fmt"
)

func main() {
    kr, err := keyring.New()
    if err != nil {
        fmt.Println(err
    }

    err := keyring.Set("mysupercoolservice", "larry", "password1")
    if err != nil 
    password, err := kr.Get("supernova", "global:LDAP")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(password)
}
```

