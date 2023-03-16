package pathutils

import (
	"strings"
	"testing"
)

func TestHomepath(t *testing.T) {
	path := Homepath()

	expecteds := []string{
		"/home/codespace",
		"/home/runner",
	}
	for _, expected := range expecteds {
		if strings.Contains(path, expected) == false {
			t.Error("Does not contain ", expected, ", Output:", path)
		}
	}
}
