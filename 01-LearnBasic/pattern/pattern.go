package pattern

import (
	"fmt"
)

/*________________
#pattern-1
*****
*****
*****
*****
*****
_________________*/

func Pattern1() {
	fmt.Printf("#pattern-1\n")
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*________________
#pattern-2
*
**
***
****
*****
_________________*/

func Pattern2() {
	fmt.Printf("#pattern-2\n")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*________________
#pattern-3
1
12
123
1234
12345
_________________*/

func Pattern3() {
	fmt.Printf("#pattern-3\n")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Printf("\n")
	}
}

/*________________
#pattern-4
1
22
333
4444
55555
_________________*/

func Pattern4() {
	fmt.Printf("#pattern-4\n")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-5
*****
****
***
**
*
_________________*/

func Pattern5() {
	fmt.Printf("#pattern-5\n")
	for i := 5; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-6
12345
1234
123
12
1
________________*/

func Pattern6() {
	fmt.Printf("#pattern-6\n")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Printf("%d", j+1)
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-7
    *
   ***
  *****
 *******
*********
______________*/

func Pattern7() {
	fmt.Printf("#pattern-7\n")
	n := 5
	for i := 1; i <= n; i++ {
		for k := n - i; k > 0; k-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-8
 *********
  *******
   *****
    ***
     *
______________*/

func Pattern8() {
	fmt.Printf("#pattern-8\n")
	n := 5
	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*n - (2 * i) + 1); j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-9
    *
   ***
  *****
 *******
*********
*********
 *******
  *****
   ***
    *
________________*/

func Pattern9() {
	fmt.Printf("#pattern-9\n")
	n := 5
	for i := 1; i <= n; i++ {
		for k := n - i; k > 0; k-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i)-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := 1; i <= n; i++ {
		for k := 1; k < i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*n - (2 * i) + 1); j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-10
*
**
***
****
*****
****
***
**
*
________________*/

func Pattern10() {
	fmt.Printf("#pattern-10\n")
	n := 5
	for i := 1; i <= (n*2)-1; i++ {

		star := i
		if i > n {
			star = 2*n - i
		}

		for j := 1; j <= star; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-11
1
01
101
0101
10101
________________*/

func Pattern11() {
	fmt.Printf("#pattern-11\n")
	n := 5
	for i := 1; i <= n; i++ {

		start := 1
		if i%2 == 0 {
			start = 0
		} else {
			start = 1
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%d", start)
			start = 1 - start
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-12
1      1
12    21
123  321
12344321
________________*/

func Pattern12() {
	fmt.Printf("#pattern-12\n")
	n := 4
	space := 2 * (n - 1)
	for i := 1; i <= n; i++ {

		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		for k := 1; k <= (2*n)-(2*i); k++ {
			fmt.Printf(" ")
		}
		//OR
		// for k := 1; k <= space; k++ {
		// 	fmt.Printf(" ")
		// }

		for l := i; l > 0; l-- {
			fmt.Printf("%d", l)
		}
		space = space - 2
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-13
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
________________*/

func Pattern13() {
	fmt.Printf("#pattern-13\n")
	n := 5
	number := 1
	for i := 1; i <= n; i++ {

		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", number)
			number = number + 1
		}
		fmt.Printf("\n")
	}
}

/*_______________
#pattern-19
**********
****##****
***####***
**######**
*########*
*########*
**######**
***####***
****##****
**********
________________*/

func Pattern19() {
	fmt.Printf("#pattern-19\n")
	n := 5
	space := 0

	for i := 0; i < n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Printf("*")
		}
		for l := 0; l < space; l++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= n-i; k++ {
			fmt.Printf("*")
		}
		space += 2
		fmt.Printf("\n")
	}

	space1 := 8
	for i := 1; i <= n; i++ {

		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		for l := 0; l < space1; l++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("*")
		}
		space1 -= 2
		fmt.Printf("\n")
	}

}

/*_______________
#pattern-20
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **
*        *
________________*/

func Pattern20() {
	fmt.Printf("#pattern-20\n")
	n := 5
	space := (2 * n) - 2
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}

		//space
		for k := 1; k <= space; k++ {
			fmt.Printf(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}

		space = space - 2
		fmt.Printf("\n")
	}

	space = 2
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Printf("*")
		}

		//space
		for k := 1; k <= space; k++ {
			fmt.Printf(" ")
		}

		for j := n - i; j > 0; j-- {
			fmt.Printf("*")
		}

		space = space + 2
		fmt.Printf("\n")
	}

}

/*_______________
#pattern-21
****
*  *
*  *
****
________________*/

func Pattern21() {
	fmt.Printf("#pattern-21\n")
	n := 4
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}

/*_______________
#pattern-22
4444444
4333334
4322234
4321234
4322234
4333334
4444444
________________*/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Pattern22() {
	fmt.Printf("#pattern-22\n")
	n := 4
	for i := 0; i < (2*n)-1; i++ {
		for j := 0; j < (2*n)-1; j++ {
			top := i
			left := j
			right := (2*n - 2) - i
			bottom := (2*n - 2) - j
			minValue := min(min(top, bottom), min(left, right))
			fmt.Printf("%d", n-minValue)
		}
		fmt.Printf("\n")
	}

}
