package math

import "fmt"

// Prime Number: A number which is exactly two factor 1 and itself
func IsPrimeNumber(n int32) {

	var i, factor int32
	var count int8 = 0
	for i = 1; i*i <= n; i++ {
		if count > 2 {
			fmt.Printf("%d Not a Prime Number\n", n)
			return
		}

		if n%i == 0 {
			count++
			if factor = n / i; factor != i {
				count++
			}
		}
	}
	if count == 2 {
		fmt.Printf("%d is a Prime Number\n", n)
		return
	}
}
