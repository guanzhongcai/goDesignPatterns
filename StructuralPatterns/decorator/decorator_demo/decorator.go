package decorator_demo

import "fmt"

// 定义一个抽象的组件
type Component interface {
	Operate()
}

// 实现一个具体的组件：组件1
type Component1 struct {
}

func (c1 *Component1) Operate() {
	fmt.Println("c1 operate")
}

// 定义一个抽象的装饰者
type Decorator interface {
	Component
	Do() // 这个是额外的方法
}

// 实现一个具体的装饰者
type Decorator1 struct {
	c Component
}

func (d1 *Decorator1) Do() {
	fmt.Println("发生了装饰行为")
}

func (d1 *Decorator1) Operate() {

	d1.Do()
	d1.c.Operate()
}
