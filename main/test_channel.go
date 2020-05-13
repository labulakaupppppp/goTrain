/**
2 * @Author: 于淼
  * @file: test_channel
3 * @Date: 2020-05-12 21:22
4 */

package main

import "fmt"

/*
通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
*/
func sum(arr []int, c chan int) {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	c <- sum

}

func cov() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	s := []int{2, 5, 1, -9, 7, 13}
	ch := make(chan int, 0)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	cov()

}
