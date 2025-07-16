package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; ; i++{
		prev := z
		z -= (z*z - x) / (2 * z)
		
		if math.Abs(z-prev) < 1e-10 {
			break
		}
		fmt.Println(i, "回目", z)
	}
	return z
}

func main() {
	fmt.Println("結果：", Sqrt(2))
	fmt.Println("math.Sqrtとの比較：", math.Sqrt(2))
}
