package channel

import "fmt"

func gen() chan int {
	in := make(chan int)
	go func() {
		for i := 2; i < 10; i++ {
			in <- i
		}
	}()
	return in
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	primes := make(chan int)

	go func() {
		in := gen()
		fmt.Printf("in is %v\n", in)
		for {
			prime := <-in
			in = filter(in, prime)
			fmt.Printf("in is %v\n", in)
			primes <- prime
		}
	}()

	return primes
}
