package channel

import (
	"fmt"
	"strconv"
	"time"
)

func boom() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	count := 0
	for {
		select {
		case <-tick:
			fmt.Printf("tick.%d\n", count)
		case <-boom:
			fmt.Printf("BOOM!%d\n", count)
			return
		default:
			fmt.Printf("    .%d\n", count)
			time.Sleep(5e7)
			count += 1
		}
	}
}

func initTimeout() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeoutChan := make(chan bool, 1)

	makeTimeout := func(tc chan bool, t int) {
		time.Sleep(time.Second * time.Duration(t))
		tc <- true
	}

	go makeTimeout(timeoutChan, 2)

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received:", msg1) // 本来说是无序的，不过还是 c1 被先执行了
	case msg2 := <-c2:
		fmt.Println("c2 received:", msg2)
	case <-timeoutChan:
		fmt.Println("Timeout, exit.")
	}
}

/*
 对于这种单条消息的 timeout 就无效了
*/
func invalidTimeout() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeoutChan := make(chan bool, 1)

	makeTimeout := func(tc chan bool, t int) {
		time.Sleep(time.Second * time.Duration(t))
		tc <- true
	}

	go makeTimeout(timeoutChan, 2)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(1e8)
			c1 <- strconv.Itoa(i)
		}
	}()

LE:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("c1 received:", msg1) // 本来说是无序的，不过还是 c1 被先执行了
		case msg2 := <-c2:
			fmt.Println("c2 received:", msg2)
		case <-timeoutChan:
			fmt.Println("Timeout, exit.")
			break LE
		}
	}
	fmt.Println("System over")
}
