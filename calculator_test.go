package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f and %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 5, want: -4},
		{a: 5, b: 1, want: 4},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f and %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := calculator.Multiply(3, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 5, want: 5},
		{a: 5, b: -2, want: -10},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f and %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 1, b: 2, want: 0.5},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.3333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(tc.want, got, 0.0001) {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Errorf("Divide by zero should return error")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 100, want: 10},
		{a: 0, want: 0},
		{a: 1, want: 1},
		{a: 10, want: 3.16227},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Errorf("Unexpected error happened, input:%f, want:%f, output:%f, error: %v", tc.a, tc.want, got, err)
		}
		if !closeEnough(tc.want, got, 0.0001) {
			t.Fatalf("Sqrt(%f), want:%f got:%f", tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
