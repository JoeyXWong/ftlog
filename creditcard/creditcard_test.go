package creditcard_test

import (
	"calculator/creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cc, err := creditcard.New("123")

	if err != nil {
		t.Fatalf(err.Error())
	}
	want := "123"

	if want != cc.GetNumber() {
		t.Errorf("want: %s, got: %s", want, cc.GetNumber())
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()

	_, err := creditcard.New("")

	if err == nil {
		t.Errorf("should return invalid number error")
	}
}
