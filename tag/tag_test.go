package tag

import "testing"

func TestNormal(t *testing.T) {
	if testing.Short() {
		t.Skip("skip normal unit test")
	}

	// to test here...
	t.Log("normal unit test without build tags")
}
