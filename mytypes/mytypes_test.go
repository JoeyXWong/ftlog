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

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", mb.Contents.String(), wantLen, gotLen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("hello")
	want := "HELLO"
	got := mb.ToUpper()
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	var x mytypes.MyInt = 12
	want := mytypes.MyInt(24)
	x.Double()
	if want != x {
		t.Errorf("want: %d, got: %d", want, x)
	}
}
