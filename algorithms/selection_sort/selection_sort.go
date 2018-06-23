package main

import (
	"fmt"

	"github.com/xiaoheigou/GootnvAlgorithms/SortAlgorithms/util"
)

func main() {
	randArr := util.RandArray(10)
	fmt.Println("befer:", randArr)
	selection_sort(randArr)
	fmt.Println("after:", randArr)
}
func selection_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		//min是当前最小值得下标
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 将遍历到的最小的放到i的位置
		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
}
