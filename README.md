# errline

Package adds file and line information to Go errors.

Errors in Go do not have stack trace or file location information attached to them.
This package only attaches short file name and line information to the error.
That way both error location information and error message can fit in one line.

The errline.Wrap function returns a new error that adds context to the original
error by recording short file name and line at the point Wrap is called.

It may be necessary to reverse the operation of errline.Wrap to retrieve the original error
for inspection. Any error value which implements this interface

```
type causer interface {
    Cause() error
}
```

can be inspected by errline.Cause.

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

