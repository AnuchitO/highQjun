//go:build integration

package fizzbuzz_test

import (
	"highQjun/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	want := "FizzBuzz"

	got := fizzbuzz.F(15)

	if got != want {
		t.Fatalf("FizzBuzz(15): expected %s, actual %s", want, got)
	}
}
