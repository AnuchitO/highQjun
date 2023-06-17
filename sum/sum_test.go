package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	r := Sum(1, 2)

	if r != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}
