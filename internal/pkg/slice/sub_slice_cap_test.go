package slice

import "testing"

func TestSlice(t *testing.T) {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	bar := foo[1:4]
	bar[1] = 99

	t.Logf("%s cap is %d, len is %d\n", "foo", cap(foo), len(foo))
	t.Logf("%s cap is %d, len is %d\n", "bar", cap(bar), len(bar))
}
