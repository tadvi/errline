package errline

import (
	"errors"
	"fmt"
)

func oops() error {
	// Do not change location of this line in the file.
	return Wrap(errors.New("the file system has gone away"))
}

func Example() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// Output: example_test.go:10: the file system has gone away
}
