package main

import (
	"flag"
	"fmt"
	"strconv"
)

var ip = flag.Int("flagname", 1234, "help message for flagname")

func main() {
	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(strconv.IntSize)
	fmt.Println("无符号0取反", ^uint(0))
	fmt.Println("有符号0取反", ^int(0))
	var a uint
	a = ^uint(0) >> 63
	fmt.Println("无符号0取反右移63位", a)
	fmt.Println("有符号0取反右移63位", ^int(0)>>63)

	// i, err := strconv.ParseInt("21474836478", 10, 32)
	// if err != nil {
	// 	panic(err)
	// }
	// println(int32(i), math.MaxInt32)
	// 会报错，int32 value out of range
	ii, err := strconv.ParseInt("A", 16, 32)
	if err != nil {
		panic(err)
	}
	println(int32(ii))

	x := ^uint32(0) // x is 0xffffffff
	i := int(x)     // i is -1 on 32-bit systems, 0xffffffff on 64-bit
	fmt.Println(i)
}
