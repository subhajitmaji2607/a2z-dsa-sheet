package math

import (
	"fmt"
)

func CountDigits(n int) {
	count := 0
	x := n
	// fmt.Printf("%d", n%10)
	// fmt.Printf("%d", n/10)
	for n != 0 {
		n = n / 10
		count++
	}
	fmt.Printf("Number of digits in %d is %d\n", x, count)
}
