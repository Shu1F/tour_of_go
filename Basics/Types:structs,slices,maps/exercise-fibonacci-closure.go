package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		tmp := a
		a = b
		b = tmp + b
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
