package fatal

import (
	"fmt"
	"os"
)

// F does Fprintf to stderr with the supplied args then calls os.Exit(1)
func F(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

// IfF passes msg and args to F() if condition is true
func IfF(condition bool, msg string, args ...any) {
	if condition {
		F(msg, args...)
	}
}

// IfErrorF calls IfF with the condition err != nil. It appends err to args
// so msg should include a final format specifier for the error
func IfErrorF(err error, msg string, args ...any) {
	IfF(err != nil, msg, append(args, err)...)
}
