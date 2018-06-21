package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	spe()
}

//监听指定信号
func spe() {
	//合建chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	//阻塞直到有信号传入
	fmt.Println("启动")
	//阻塞直至有信号传入
	s := <-c
	fmt.Println("退出信号", s)
}

//监听所有的信号
func all() {
	c := make(chan os.Signal)
	signal.Notify(c)

	fmt.Println("启动")
	s := <-c
	//阻塞，直到接收到退出信号
	fmt.Println("退出信号", s)
}
