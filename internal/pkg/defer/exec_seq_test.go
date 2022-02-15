/*
多个 defer 出现的时候，它是一个 “栈” 的关系，也就是先进后出
*/
package defertest

import (
	"fmt"
	"testing"
)

func func1() {
    fmt.Println("A")
}

func func2() {
    fmt.Println("B")
}

func func3() {
    fmt.Println("C")
}

func TestExecSeq(t *testing.T) {
    defer func1()
    defer func2()
    defer func3()
}

/* Output:
C
B
A
*/