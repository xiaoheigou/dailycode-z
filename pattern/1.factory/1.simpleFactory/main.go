package main

import "fmt"

type Operater interface {
	//运算接口
	Operate(int, int) int
}
type AddOperater struct {
}

func (this *AddOperater) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type SubOperater struct {
}

func (this *SubOperater) Operate(rhs int, lhs int) int {
	return rhs - lhs
}

type factory struct {
}

func (this *factory) CreateOperate(operatename string) Operater {
	switch operatename {
	case "+":
		return &AddOperater{}
	case "-":
		return &SubOperater{}
	default:
		panic("无效的运算符")
		return nil
	}
}
func CreateFactory() *factory {

	return &factory{}
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Operator := CreateFactory().CreateOperate("+")
	fmt.Printf("add result is %d\n", Operator.Operate(2, 4))
}
