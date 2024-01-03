package math

import "fmt"

func PalindromeNumber(n int) {
	if reverse := ReverseNumber(n); reverse == n {
		fmt.Printf("Palindrome Number\n")
	} else {
		fmt.Printf("Not Palindrome Number\n")
	}
}
