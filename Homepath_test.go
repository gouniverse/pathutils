package pathutils

import (
	"strings"
	"testing"
)

func TestHomepath(t *testing.T) {
	path := Homepath()

	expected := "/home/codespace"
	if strings.Contains(path, expected) == false {
		t.Error("Does not contain ", expected, ", Output:", path)
	}
}
