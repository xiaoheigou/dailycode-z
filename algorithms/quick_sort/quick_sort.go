package main

import (
	"fmt"

	"github.com/xiaoheigou/GootnvAlgorithms/SortAlgorithms/util"
)

// quick sort ref:http://blog.csdn.net/morewindows/article/details/6684558
func main() {
	randArr := util.RandArray(10)
	fmt.Println("befer:", randArr)
	quickSort(randArr, 0, len(randArr))
	fmt.Println("befer:", randArr)
}

func quickSort(arr []int, left, right int) {
	temp := arr[left]
	p := left

	i, j := left, right
	for i <= j {
		//从右边开始找，找到第一个比temp小的数
		for j >= p && arr[j] >= temp {
			j--
		}
		//把[j]填到[p]的位置.此算法默认p永远就是那个坑
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		//从左向右找到比temp大的第一个数
		for i <= p && arr[i] <= temp {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}

	}
	arr[p] = temp //左右反复扫描，temp终于找到自己的萝卜坑了

	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}

	return
}
