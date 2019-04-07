package main

import "fmt"

// 鸭子基类
type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) Swim() {
	fmt.Println("I'm swimming")
}

func (d *Duck) Display() {
	fmt.Println("I'm Displaying")
}

func (d *Duck) SetFly(flyBehavior FlyBehavior) {
	d.flyBehavior = flyBehavior
}

func (d *Duck) SetQuack(quackBehavior QuackBehavior) {
	d.quackBehavior = quackBehavior
}

func (d *Duck) PeformFly() {
	if d.flyBehavior != nil {
		d.flyBehavior.Fly()
	}
}

func (d *Duck) PeformQuack() {
	if d.quackBehavior != nil {
		d.quackBehavior.Quack()
	}
}

// 飞行接口
type FlyBehavior interface {
	Fly()
}

// 飞行算子
type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings.")
}

// 飞行算子
type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	fmt.Println("I can not fly.")
}

// 呱呱叫接口
type QuackBehavior interface {
	Quack()
}

// 呱呱叫算子
type Quack struct{}

func (f *Quack) Quack() {
	fmt.Println("I'm Quacking.")
}

// 呱呱叫算子
type Squeak struct{}

func (f *Squeak) Quack() {
	fmt.Println("I'm Squeaking.")
}

// 呱呱叫算子
type QuackNoWay struct{}

func (f *QuackNoWay) Quack() {
	fmt.Println("I can not Quack.")
}

// 工厂方法
func NewMallardDuck() *Duck {
	return &Duck{
		flyBehavior:   &FlyWithWings{},
		quackBehavior: &Squeak{},
	}
}

func NewModelDuck() *Duck {
	return &Duck{
		flyBehavior:   &FlyNoWay{},
		quackBehavior: &QuackNoWay{},
	}
}

func main() {
	var mallard *Duck = NewMallardDuck()
	mallard.PeformFly()
	mallard.PeformQuack()
	mallard.SetFly(&FlyNoWay{})
	mallard.PeformFly()

	model := NewModelDuck()
	model.Display()
	mallard.PeformFly()
	mallard.PeformQuack()
	mallard.SetFly(&FlyWithWings{})
	mallard.PeformFly()
}
