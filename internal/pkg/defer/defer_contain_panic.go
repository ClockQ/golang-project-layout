/*
defer 中包含 panic
panic 仅有最后一个可以被 revover 捕获
触发 `panic("panic")` 后 defer 顺序出栈执行，第一个被执行的 defer 中 会有 `panic("defer panic")` 异常语句，
这个异常将会覆盖掉 main 中的异常 `panic("panic")`，最后这个异常被第二个执行的 defer 捕获到。
*/
package defertest

import (
	"fmt"
	"testing"
)

func TestDeferContainPanic(t *testing.T) {
    defer func() {
        if err := recover(); err != nil{
            fmt.Println(err)
        }else {
            fmt.Println("fatal")
        }
     }()
 
     defer func() {
         panic("defer panic")
     }()
 
     panic("panic")
}

/* Output:
defer panic
*/