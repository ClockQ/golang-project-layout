package main

import "fmt"

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	bar := foo[1:4]
	bar[1] = 99

	fmt.Printf("%s cap is %d, len is %d\n", "foo", cap(foo), len(foo))
	fmt.Printf("%s cap is %d, len is %d\n", "bar", cap(bar), len(bar))
}
