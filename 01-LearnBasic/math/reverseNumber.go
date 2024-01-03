package math

import "fmt"

/*
How to extract digits from the end of a number?

To extract the last digit, if you divide a number by 10,
then the remainder will be the last digit. We can simply use the modulo(%) operator to do this,
for example the last digit of 123 will be (123 % 10), which is 3.

To reduce the number by one digit from the end, simply divide the number by 10.
example: to reduce 123 to 12, simply do (123/10) which is equal to 12.

To create a number from digits: The idea is to start with 0, and for every digit,
multiply the number generated so far by 10, and add the digit to it.

For example, to create a number from digits: [1,2,3]:

Consider the number, num = 0.
*/
func ReverseNumber(n int) int {
	reverse := 0
	x := n
	for n != 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10
	}

	fmt.Printf("The reverse of the %d is %d\n", x, reverse)
	return reverse
}
