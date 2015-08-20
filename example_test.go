package keyring_test

import (
	"fmt"
	"github.com/thinmintaddict/keys"
	"os"
)

func ExampleMain() {
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
	// Output:
	// larry's password for myservice is 'password1'.
}
