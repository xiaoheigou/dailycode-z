package main

import "fmt"

//简单工厂模式是通过传递不同的参数生成不同的实例，缺点就是扩展不同的类别时需要修改代码。

//工厂方法模式为每一个product提供一个工程类，通过不同工厂创建不同实例。
// 相当于给一个工厂加了流水线，用来生产同一商品不同型号
func main() {
	var factoryInterface AnimalFactory

	factoryInterface = &BirdFactory{}
	animal := factoryInterface.CreateAnimal()
	animal.move()

}

type action interface {
	move()
}
type Bird struct {
	name string
}
type Fish struct {
	name string
}

func (this *Bird) move() {
	fmt.Println("this is a bird")
}
func (this *Fish) move() {
	fmt.Println("this is a fish")
}

type AnimalFactory interface {
	CreateAnimal() action
}
type BirdFactory struct {
}

func (this *BirdFactory) CreateAnimal() action {
	return &Bird{}
}

type FishFactory struct {
}

func (this *FishFactory) CreateAnimal() action {
	return &Fish{}
}
