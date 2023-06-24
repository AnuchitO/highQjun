package change_test

import (
	"highQjun/change"
	"testing"
)

func TestSumUsage(t *testing.T) {
	want := 4

	got := change.Sum([]int{1, 1, 1, 1})

	if got != want {
		t.Fatalf("sum([]int{1, 1, 1, 1}): expected %d, actual %d", want, got)
	}
}
