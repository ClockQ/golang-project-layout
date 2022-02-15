/*
我们知道，能够触发 defer 的是遇见 return (或函数体到末尾) 和遇见 panic。
*/
package defertest

import (
	"fmt"
	"testing"
)

/*
A. defer 遇见 panic，但是并不捕获异常的情况
*/
func defer_call() {
    defer func() { fmt.Println("defer: panic 之前1") }()
    defer func() { fmt.Println("defer: panic 之前2") }()

    panic("异常内容")  //触发defer出栈

    defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}

func TestDeferPanic1(t *testing.T) {
    defer_call()

    fmt.Println("main 正常结束")
    /* Output:
    defer: panic 之前2
    defer: panic 之前1
    panic: 异常内容
    //... 异常堆栈信息
    */
}

/*
B. defer 遇见 panic，并捕获异常
defer 最大的功能是 panic 后依然有效
所以 defer 可以保证你的一些资源一定会被关闭，从而避免一些异常出现的问题
*/
func defer_call2() {

    defer func() {
        fmt.Println("defer: panic 之前1, 捕获异常")
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    defer func() { fmt.Println("defer: panic 之前2, 不捕获") }()

    panic("异常内容")  //触发defer出栈

    defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func TestDeferPanic2(t *testing.T) {
    defer_call2()

    fmt.Println("main 正常结束")
    /* Output:
    defer: panic 之前2, 不捕获
    defer: panic 之前1, 捕获异常
    异常内容 //捕获处理异常，打印异常信息了
    main 正常结束
    */
}