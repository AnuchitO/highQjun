package isexample

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/matryer/is"
)

func ParseBinary(b string) (bool, error) {
	return true, nil
}

func Test(t *testing.T) {
	ss := is.New(t)

	b, err := ParseBinary("0")

	ss.NoErr(err)
	ss.Equal(b, true)
	ss.Equal([]string{"a", "b", "c"}, []string{"a", "b", "c"})
	ss.True(true)

	pretty.Println([]string{"a", "b", "c"})
}
