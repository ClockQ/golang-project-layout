/*
在没有 defer 的情况下，其实函数的返回就是与 return 一致的，但是有了 defer 就不一样了。
我们知道，先 return，再 defer，所以在执行完 return 之后，还要再执行 defer 里的语句，依然可以修改本应该返回的结果。
*/
package defertest

import (
	"fmt"
	"testing"
)

func returnButDefer() (t int) {  //t初始化0， 并且作用域为该函数全域
    defer func() {
        t = t * 10
    }()
    return 1
}

func TestReturnInitDefer(t *testing.T) {
    fmt.Println(returnButDefer())
}

/* Output:
10
*/