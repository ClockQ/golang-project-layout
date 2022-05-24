package channel

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
