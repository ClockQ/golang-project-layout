/*
return 之后的语句先执行，defer 后的语句后执行
*/
package defertest

import (
	"fmt"
	"testing"
)

func deferFunc() int {
    fmt.Println("defer func called")
    return 0
}

func returnFunc() int {
    fmt.Println("return func called")
    return 0
}

func returnAndDefer() int {

    defer deferFunc()

    return returnFunc()
}

func TestDeferAndReturn(t *testing.T) {
    returnAndDefer()
}

/* Output:
return func called
defer func called
*/