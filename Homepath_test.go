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

	expected := strings.Join(expecteds, ",")

	if strings.Contains(expected, path) == false {
		t.Error("Does not contain ", path, ", Output:", expected)
	}
}
