package main

import (
	"fmt"
	"math"
)

func main() {
	findMedianSortedArrays([]int{1, 2}, []int{3, 4})
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	//这样子得到的n1是数组长度较小的值
	// k 是我们一共取了多少个数
	k := (n1 + n2 + 1) / 2

	fmt.Println("k的值是：", k)
	l, r := 0, n1
	m1, m2 := 0, 0
	// 开始循环
	for {
		m1 = (l + r) / 2
		m2 = k - m1
		if l >= r {
			fmt.Println("r--l:", r, l)
			break
		}
		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}
	fmt.Println("退出当前循环，当前m1,m2:", m1, m2)
	t1, t2 := 0, 0
	if m1 <= 0 {
		t1 = -1
	} else {
		t1 = nums1[m1-1]
	}

	if m2 <= 0 {
		t2 = -1
	} else {
		t2 = nums2[m2-1]
	}
	c1 := math.Max(float64(t1), float64(t2))
	if (n1+n2)%2 == 1 {
		return c1
	}
	if m1 >= n1 {
		t1 = 999999999
	} else {
		t1 = nums1[m1]
	}
	if m2 >= n2 {
		t2 = 9999999999
	} else {
		t2 = nums2[m2]
	}
	fmt.Println("t1--t2", t1, t2)
	c2 := math.Min(float64(t1), float64(t2))
	fmt.Println("c1---c2", c1, c2)
	return (c1 + c2) * 0.5

}
