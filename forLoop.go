package main

import "fmt"

func main() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 34, 5}

	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	for a < b {
		a++
		fmt.Println(a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	var i int
	var j int

	for i = 2; i < 100; i++ {
		for j = 2; j < (i/j); j++ {
			if j%j == 0 {
				break

			}
		}
		if (j > (i / j)) {
			fmt.Printf("%d  是素数\n", i)
		}
	}

	for true {
		fmt.Println("无限循环")
	}
}
