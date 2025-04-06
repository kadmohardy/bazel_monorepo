package library

import (
	"testing"
)

// TestToUppercase checks that returns a uppercase string
func TestToUppercase(t *testing.T) {
	actual := ToUppercase("test")
	if actual != "TEST" {
		t.Errorf("Strings diferentes: esperado %q, mas veio %q", "TEST", actual)
	}
}
