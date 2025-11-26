package Task1

import (
	"fmt"
	"sort"
)

func Run() {

	//检查数组没有重复的数
	a := []int{1, 2, 3, 4, 2, 4, 56, 8, 7, 6, 3, 4, 99}
	b, ok := onlyOne(a).([]int)
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

	//检查最长的公共前缀
	strs := []string{"flower", "flow", "flight"} // 期待结果为 "fl":ml-citation{ref="3,4" data="citationList"}
	result1 := longestCommonPrefix(strs)
	fmt.Printf("最长的公共前缀是：%s", result1)

	//加一
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))

	//删除重复的，返回唯一的数量
	numbers := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("origin a: %v, ", numbers)
	length := removeRepeat(&numbers)
	fmt.Printf("new a: %v, len: %d\n", numbers, length)

	//合并区间
	testMerge := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {0, 0}},
	}
	for i, test := range testMerge {
		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: %v\n", test)
		fmt.Printf("输出: %v\n\n", margedIntervals(test))
	}

	//检查切片是否满足两和等于目标，返回他们索引
	fmt.Printf("return: %v\n", findSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("return: %v\n", findSum([]int{3, 2, 4}, 6))
	fmt.Printf("return: %v\n", findSum([]int{3, 3}, 6))

}

// 判断没有重复的数返回切片数组
func onlyOne(in []int) interface{} {
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
	var stack []rune
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

// 返回这些字符串的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 {
			if len(strs[i]) >= len(prefix) && strs[i][:len(prefix)] == prefix {
				break
			} else {
				prefix = prefix[:len(prefix)-1]
			}
		}

		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

func plusOne(a []int) []int {
	length := len(a)
	result := make([]int, length+1)
	step := 0
	for i := len(result) - 1; i >= 0; i-- {
		j := i - 1
		if i == length {
			step = a[j] + 1
			result[i] = step % 10
		} else if i != 0 {
			step = a[j] + step/10
			result[i] = step % 10
		} else {
			if step == 10 {
				result[0] = step / 10
			}
		}

	}
	if result[0] == 0 {
		result = result[1:]
	}
	return result
}

func removeRepeat(s *[]int) int {
	nonRepeat, cursor := 0, 1
	for i := 0; i < len(*s); i++ {
		cursor = i
		if (*s)[nonRepeat] != (*s)[cursor] {
			nonRepeat++
			(*s)[nonRepeat] = (*s)[cursor]
		}
	}
	if nonRepeat != len(*s)-1 {
		*s = (*s)[:nonRepeat+1]
	}
	return nonRepeat + 1
}

// 合并区间
func margedIntervals(s [][]int) [][]int {
	if s == nil || len(s) == 0 {
		return nil
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})

	merged := [][]int{s[0]}
	for i := 0; i < len(s); i++ {
		current := s[i]
		lastMerged := merged[len(merged)-1]
		//当前起始值<= 最后合并区结束值
		if current[0] <= lastMerged[len(lastMerged)-1] {
			//存在重叠，取max作为结束
			if current[len(current)-1] > lastMerged[len(lastMerged)-1] {
				lastMerged[len(lastMerged)-1] = current[len(current)-1]
			}
		} else {
			//没有重叠区域添加新区
			merged = append(merged, current)
		}
	}
	return merged
}

func findSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) < 2 {
		return []int{}
	}
	numToIndex := make(map[int]int)
	for i, v := range numbers {
		complement := target - v
		if _, exists := numToIndex[complement]; exists {
			return []int{numToIndex[complement], i}
		} else {
			numToIndex[v] = i
		}
	}

	return []int{}
}
