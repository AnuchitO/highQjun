package benchmark

import "strings"

func Cat(xs []string) string {
	var b strings.Builder
	b.Grow(len(xs) * 8)
	b.WriteString(xs[0])
	for _, v := range xs[1:] {
		b.WriteString(" ")
		b.WriteString(v)
	}
	return b.String()
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
