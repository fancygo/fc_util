package fc_sys

import (
	"testing"
)

func TestIf(t *testing.T) {
	a, b := 3, 4
	c := If(a > b, a, b).(int)
	println(c)
}
