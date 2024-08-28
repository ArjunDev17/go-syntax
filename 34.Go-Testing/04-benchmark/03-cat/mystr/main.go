package mystr

import "strings"

func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
func Cat1(xs []string) string {
	var sb strings.Builder
	for i, v := range xs {
		sb.WriteString(v)
		if i < len(xs)-1 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
