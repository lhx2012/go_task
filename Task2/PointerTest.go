package Task2

import (
	"fmt"
)

func pointAddNum(num *int) {
	*num += 10
}

func pointDouble(nums *[]int) {
	for i, _ := range *nums {
		(*nums)[i] *= 2
	}
}

func pointRun() {
	num := 50
	fmt.Printf("Num初始值：%d\n", num)
	pointAddNum(&num)
	fmt.Printf("Num处理过后的值:%d\n", num)

	nums := []int{2, 5, 7, 9, 16}
	fmt.Printf("Nums初始值：%v\n", nums)
	pointDouble(&nums)
	fmt.Printf("调用Double后的值：%v\n", nums)
}
