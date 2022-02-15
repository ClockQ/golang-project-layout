/*
那么这 4 个函数的先后执行顺序是什么呢？这里面有两个 defer， 所以 defer 一共会压栈两次，先进栈 1，后进栈 2。 那么在压栈 function1 的时候，需要连同函数地址、函数形参一同进栈，那么为了得到 function1 的第二个参数的结果，所以就需要先执行 function3 将第二个参数算出，那么 function3 就被第一个执行。同理压栈 function2，就需要执行 function4 算出 function2 第二个参数的值。然后函数结束，先出栈 fuction2、再出栈 function1.

所以顺序如下：
- defer 压栈 function1，压栈函数地址、形参 1、形参 2 (调用 function3) –> 打印 3
- defer 压栈 function2，压栈函数地址、形参 1、形参 2 (调用 function4) –> 打印 4
- defer 出栈 function2, 调用 function2 –> 打印 2
- defer 出栈 function1, 调用 function1–> 打印 1
*/
package defertest

import (
	"fmt"
	"testing"
)

func function(index int, value int) int {
    fmt.Println(index)
    return index
}

func TestDeferFuncContainFunc(t *testing.T) {
    defer function(1, function(3, 0))
    defer function(2, function(4, 0))
}

/* Output:
3
4
2
1
*/