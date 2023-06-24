package change

import "testing"

func TestSum(t *testing.T) {
	want := 4
	got := Sum([]int{1, 1, 1, 1})

	if got != want {
		t.Fatalf("sum([]int{1, 1, 1, 1}): expected %d, actual %d", want, got)
	}
}
