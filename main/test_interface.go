/**
2 * @Author: 于淼
  * @file: test_interface
3 * @Date: 2020-05-12 20:17
4 */

package main

import "fmt"

//接口
type Animal interface {
	call()
	run()
}

//结构体 也就是类
type Cat struct {
	name string
	miao string
}

type Dog struct {
	name string
	wang string
}

//接口的实现
func (cat Cat) call() {
	fmt.Printf("%s is calling %s~\n", cat.name, cat.miao)
}

func (cat *Cat) run() {
	fmt.Printf("%s is running %s~\n", cat.name, cat.miao)
}

func (dog Dog) call() {
	fmt.Printf("%s is calling %s~\n", dog.name, dog.wang)
}
func (dog Dog) run() {
	panic("implement me")
}

func main() {
	var animal Animal
	//animal = new(Cat)
	animal = &Cat{"猫", "miao"}
	animal.call()
	animal.run()

	//animal=new(Dog)
	animal = Dog{"狗", "wang"}
	animal.call()
}
