package pathutils

import (
	"os"
	"path/filepath"
	"runtime"
)

// Gopath returns the value of the $GOPATH environment variable or its default
// value if not set.
func Gopath() string {
	if r := os.Getenv("GOPATH"); r != "" {
		return r
	}

	switch runtime.GOOS {
	case "plan9":
		return os.Getenv("home")
	case "windows":
		return filepath.Join(os.Getenv("USERPROFILE"), "go")
	default:
		return filepath.Join(os.Getenv("HOME"), "go")
	}
}
