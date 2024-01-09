package math

import "fmt"

func PalindromeNumber(n int) {
	if n < 0 {
		fmt.Printf("Not Palindrome Number\n")
	}

	if reverse := ReverseNumber(n); n < 0 || reverse == n {
		fmt.Printf("Palindrome Number\n")
	} else {
		fmt.Printf("Not Palindrome Number\n")
	}
}
