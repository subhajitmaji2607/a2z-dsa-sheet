package math

import (
	"fmt"
	"math"
)

func PrintAllDivisors(n int32) {
	//brute force approach: time complexity O(n);
	var i int32
	for i = 1; i <= n; i++ {
		if reminder := (n % i); reminder == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")

	//Optimize approach
	sqrt := int32(math.Sqrt(float64(n)))

	//**Sqrt is an mathematical function so it might some additional time to generate
	// the result to reduce this instead of calculating sqrt we can write (i*i <= n)
	var j int32
	for j = 1; j <= sqrt; j++ {
		if reminder := (n % j); reminder == 0 {
			fmt.Printf("%d ", j)
			var factor int32 = n / j
			if factor != j {
				fmt.Printf("%d ", factor)
			}
		}
	}
	// for j = 1; j*j <= n; j++ {
	// 	if reminder := (n % j); reminder == 0 {
	// 		fmt.Printf("%d ", j)
	// 		var factor int32 = n / j
	// 		if factor != j {
	// 			fmt.Printf("%d ", factor)
	// 		}
	// 	}
	// }
	fmt.Printf("\n")
}
