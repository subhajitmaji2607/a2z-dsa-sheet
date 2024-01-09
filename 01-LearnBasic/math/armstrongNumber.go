package math

import (
	"fmt"
	"strconv"
)

func powerRoot(n, power int) int {
	var number int = 1
	for i := 1; i <= power; i++ {
		number = number * n
	}
	return number
}

func ArmstrongNumber(n int) {
	originalNumber := n
	power := len(strconv.Itoa(n))
	armstrongNumber := 0
	for n != 0 {
		x := n % 10
		armstrongNumber = armstrongNumber + powerRoot(x, power)
		n = n / 10
	}

	if originalNumber == armstrongNumber {
		fmt.Printf("Yes, it is an Armstrong Number\n")
	} else {
		fmt.Printf("No, it is not an Armstrong Number\n")
	}
}
