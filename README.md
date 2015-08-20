# Keys

Keys is a cross-platform keyring library for golang, utilizing gnome-kerying onLinux or the Mac OSX Keychain to Get, Set, or Delete passwords.

### Dependencies
- libglib2.0-dev (gnome-keyring support)

### Example Usage
```go
	// using the New() function returns a struct with Get(), Set(), and Delete() as methods.
	kr, err := keyring.New()
	if err != nil {
		fmt.Println(err)
	}

	// Set() takes (service, username, password) as arguments
	err = kr.Set("myservice", "larry", "password1")
	if err != nil {
		fmt.Println(err)
	}
	if kr.IsAvailable() != true {
		fmt.Println("Keyring not available.")
		os.Exit(1)
	}
	// Get() and Delete() takes (service, username).
	password, err := kr.Get("myservice", "larry")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("larry's password for myservice is '%s'.\n", password)

	err = kr.Delete("myservice", "larry")
	if err != nil {
		fmt.Println(err)
	}
```

### Testing
```bash
~ # go get github.com/rakyll/keys
~ # go test $GOPATH/src/github.com/rakyll/keys
```
