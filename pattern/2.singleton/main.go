package main

import (
	"fmt"
	"sync"
)

//Singleton 是单例模式类
type Singleton struct {
	count int
}

func (s *Singleton) save() {
	s.count++
	fmt.Println(s.count)

}

var singleton *Singleton
var once sync.Once

func main() {
	sin := GetInstance()
	sin.save()
	for i := range []int{1, 2, 3, 4, 4, 5} {
		go func(i int) {
			p := GetInstance()
			fmt.Print(i, "--->")
			p.save()
			fmt.Println()

		}(i)
	}
	for i := range []int{1, 2, 3, 4, 4, 5} {
		go func(i int) {
			p := GetInstance()
			fmt.Print(i, "--->")
			p.save()
			fmt.Println()

		}(i)
	}

	select {}

}

// once 保证这个方法只执行一次
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
