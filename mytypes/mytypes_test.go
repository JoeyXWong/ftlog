package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()

	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)

	got := input.Twice()

	if got != want {
		t.Errorf("want: %d got: %d", want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()

	input := mytypes.MyString("testing")
	want := 7

	got := input.StringLen()

	if got != want {
		t.Errorf("input:%s, want:%d, got:%d", input, want, got)
	}
}
