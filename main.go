package main

import "fmt"

func main() {

	//检查数组没有重复的数
	a := []int{1, 2, 3, 4, 2, 4, 56, 8, 7, 6, 3, 4, 99}
	b, ok := OnlyOne(a).([]int)
	if ok {
		for i, v := range b {
			fmt.Printf("数组存在唯一数索引=%d值为:%d\n", i, v)
		}
	}

	//检查回文数
	testNumbers := []int{121, -121, 10, 0, 1221, 12321}
	for _, num := range testNumbers {
		result := isPalindrome(num)
		fmt.Printf("数字 %d 是回文数吗？ %t\n", num, result)
	}

	//检查符号匹配
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[()]}",
		"", // 空字符串
	}
	for _, testStr := range testCases {
		result := isValidString(testStr)
		fmt.Printf("输入: %-10s -> 有效: %-5t\n", "\""+testStr+"\"", result)
	}

}

// 判断没有重复的数返回切片数组
func OnlyOne(in []int) interface{} {
	if in == nil || len(in) == 0 {
		return nil
	}

	temp := make(map[int]int, len(in))

	for _, v := range in {
		value, exist := temp[v]
		if exist {
			temp[v] = value + 1
		} else {
			temp[v] = 1
		}
	}

	result := make([]int, 0, len(temp))
	for k, v := range temp {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

// 判断一个整数是否是回文数
func isPalindrome(x int) bool {
	// 负数不是回文数
	// 如果数字最后一位是0，那么只有数字本身是0才满足回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	// 只反转数字的一半进行比较
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，通过 revertedNumber/10 去掉中间位
	return x == revertedNumber || x == revertedNumber/10
}

// 判断字符串是否符号匹配
func isValidString(x string) bool {

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, value := range x {

		switch value {
		case '(', '{', '[':
			stack = append(stack, value)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != pairs[value] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
