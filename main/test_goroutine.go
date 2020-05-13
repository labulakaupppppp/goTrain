/**
2 * @Author: 于淼
  * @file: test_goroutine
3 * @Date: 2020-05-12 21:15
4 */

package main

import (
	"fmt"
	"time"
)

//轻量级线程

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
输出结果：
hello
world
hello
world
hello
world
world
hello
hello
world

没有固定的先后顺序；因为它们是两个 goroutine 在执行
*/
