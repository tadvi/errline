# errline

Package adds file and line information to Go errors.

Errors in Go do not have stack trace or file location information attached to them.
Stack traces provide extensive location failure information but are very verbose.
This package only attaches short file name and line information to the errors.

Usage:

1. Do not use it inside generic packages. Your libraries should simply return errors. That's it.
2. Use it in application level packages and application code to capture failure point information.

Example:

// Your application library packages or main package.

```
func SomeWork() error {
	// Here we call some library function.
	if err := ftp.Connect(addr); err != nil {
		// Here we wrap error with file and line information.
		// If error is already wrapped with file information then Wrap returns original error.
		return errline.Wrap(err);
	}
}

func main() {
	if err := SomeWork; err != nil {
		// IMPORTANT: only +v verb will print short file name and line number.
		// Other verbs simply print err.Error() without file information.
		log.Printf("%+v", err)  
	}
}

```

