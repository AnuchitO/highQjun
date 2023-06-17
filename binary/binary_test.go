package binary

import "testing"

func TestParseBinary(t *testing.T) {
	t.Run("should return 1 when input 01", func(t *testing.T) {
		r, _ := ParseBinary("01")

		if r != 1 {
			t.Errorf("ParseBinary(\"01\") failed: expected %d, got %d", 1, r)
		}
	})

	t.Run("should return 2 when input 10", func(t *testing.T) {
		r, _ := ParseBinary("10")

		if r != 2 {
			t.Errorf("ParseBinary(\"10\") failed: expected %d, got %d", 1, r)
		}
	})
}
