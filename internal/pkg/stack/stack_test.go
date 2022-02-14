package stack

import "testing"

var st1 Stack

func TestStack(t *testing.T) {
	st1.Push("Brown")
	st1.Push(3.14)
	st1.Push(100)
	st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		t.Log(item)
	}
}

/* Output:
[Java C++ Python C# Ruby]
100
3.14
Brown
*/