package version

import (
	"testing"
)

func TestVersionString(t *testing.T) {
	Version = "4.3.6"

	expected := "crafty 4.3.6"
	actual := String()

	if actual != expected {
		t.Errorf("Version does not match. expected: %q, actual: %q", expected, actual)
	}
}
