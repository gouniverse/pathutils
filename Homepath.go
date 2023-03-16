package pathutils

import (
	"os"
	"runtime"
)

// Homepath returns the user's home directory path.
func Homepath() string {
	// go1.8: https://github.com/golang/go/blob/74628a8b9f102bddd5078ee426efe0fd57033115/doc/code.html#L122
	switch runtime.GOOS {
	case "plan9":
		return os.Getenv("home")
	case "windows":
		return os.Getenv("USERPROFILE")
	default:
		return os.Getenv("HOME")
	}
}
