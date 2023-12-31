package binary

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidDigit = errors.New("invalid digit")

/*
	ParseBinary

input: a string of binary digits
  - "11", "1001", "11010", "10001101000"

output: the decimal equivalent of the binary digits or an error

"11" -> 3
"1001" -> 9
"11010" -> 26
"10001101000" -> 1128
*/
func ParseBinary(bin string) (int, error) {
	dec := 0
	for _, v := range bin {
		n, ok := strconv.Atoi(string(v))
		switch {
		case ok != nil:
			msg := fmt.Sprintf("\"%c\" is not a vaid digit", v)
			return 0, errors.New(msg)
		case n < 0 || 1 < n:
			msg := fmt.Sprintf("\"%d\" is not a binary digit", n)
			return 0, errors.New(msg)
		default:
			dec = dec*2 + n
		}
	}
	return dec, nil
}
