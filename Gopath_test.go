package pathutils

import (
	"strings"
	"testing"
)

func TestGopath(t *testing.T) {
	path := Gopath()

	expected := "/go"
	if strings.Contains(path, expected) == false {
		t.Error("Does not contain ", expected, ", Output:", path)
	}
}
