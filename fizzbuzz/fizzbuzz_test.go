package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{9, "Fizz"},
		{10, "Buzz"},
		{12, "Fizz"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		if got := F(test.input); got != test.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", test.input, got, test.want)
		}
	}
}
