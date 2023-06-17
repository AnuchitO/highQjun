package binary

import "testing"

var tests = []struct {
	name  string
	input string
	want  int
}{
	{"#01", "1", 1},
	{"#10", "10", 2},
	{"#11", "11", 3},
	{"#100", "100", 4},
	{"#1001", "1001", 9},
	{"#11010", "11010", 26},
	{"#10001101000", "10001101000", 1128},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, _ := ParseBinary(tt.input)

			// Assert
			if got != tt.want {
				t.Errorf("want %d,but got %d", tt.want, got)
			}
		})
	}
}

func TestParseBinaryError(t *testing.T) {
	_, err := ParseBinary("abc")

	if err.Error() != "\"a\" is not a vaid digit" {
		t.Errorf("want %s,but got %s", "\"a\" is not a vaid digit", err.Error())
	}
}

func TestParseBinaryErrorNotBinary(t *testing.T) {
	_, err := ParseBinary("3")

	if err.Error() != "\"3\" is not a binary digit" {
		t.Errorf("want %s,but got %s", "\"a\" is not a vaid digit", err.Error())
	}
}
