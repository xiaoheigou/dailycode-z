package main

import (
	"fmt"

	"github.com/xiaoheigou/GootnvAlgorithms/SortAlgorithms/util"
)

func main() {
	randArr := util.RandArray(10)
	fmt.Println("befer:", randArr)
	bubble_sort(randArr)
	fmt.Println("after:", randArr)

}
func bubble_sort(Randarray []int) {
	for i := 0; i < len(Randarray); i++ {
		for j := 0; j < len(Randarray)-1-i; j++ {
			if Randarray[j] > Randarray[j+1] {
				tmp := Randarray[j]
				Randarray[j] = Randarray[j+1]
				Randarray[j+1] = tmp
			}
		}
	}

}
