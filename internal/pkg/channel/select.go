package channel

import (
	"fmt"
)

func randomProcess() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "world"
	c1 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received:", msg1) // 本来说是无序的，不过还是 c1 被先执行了
	case msg2 := <-c2:
		fmt.Println("c2 received:", msg2)
	default:
	}
}
