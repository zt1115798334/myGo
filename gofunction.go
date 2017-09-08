package main

import "fmt"

func main() {
	//var a, b int = 1, 2
	//var maxNum int
	//maxNum = max(a, b)
	//fmt.Println(maxNum)
	//
	//var st1, st2 string = swap("xiaomi", "魅族")
	//fmt.Println(st1, st2)

	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swap3(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	var ip *int = &a
	fmt.Println(*ip)

}

func max(num1, num2 int) int {
	var ret int
	if num1 > num2 {
		ret = num1
	}
	if num1 < num2 {
		ret = num2
	}
	return ret
}

func swap(st1, st2 string) (string, string) {
	return st1, st2
}

func swap2(num1, num2 int) int {
	var temp int
	temp = num1
	num1 = num2
	num2 = temp
	return temp
}

func swap3(num1 *int, num2 *int) int {
	var temp int
	temp = *num1
	*num1 = *num2
	*num2 = temp
	return temp
}
