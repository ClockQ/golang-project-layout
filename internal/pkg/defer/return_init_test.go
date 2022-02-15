/*
只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为 0，而且在函数体作用域可见
*/
package defertest

import (
	"fmt"
	"testing"
)

func DeferFunc(i int) (t int) {
    fmt.Println("t = ", t)
    return 2
}

func TestReturnInit(t *testing.T) {
    DeferFunc(10)
}

/* Output:
t =  0
*/