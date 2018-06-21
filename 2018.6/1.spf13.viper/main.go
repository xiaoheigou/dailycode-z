package main

import (
	"fmt"

	"github.com/xiaoheigou/x/2018.6/1.spf13.viper/config"
)

var sid string

func main() {
	config.Setconfig()
}
func init() {
	sid = "hahah"
	fmt.Println("main init")
}
