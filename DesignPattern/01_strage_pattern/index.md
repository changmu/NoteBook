### 背景
1. 你需要中途接手开发一个模拟鸭子的游戏，游戏中有各种鸭子，它们有着不同的外貌，会游泳swim，会叫quack，后面可能会添加其他行为，比如飞行fly。
2. 你的上一任开发已经离职，他对此系统采用了标准的OO技术，设计了一个鸭子超类，并让各种鸭子继承此超类。

### 问题
1. 代码重复问题。各种鸭子叫的方式不尽相同，如果每种鸭子都实现一份swim方法的话会出现大量重复代码。
2. 扩展性问题。如果给父类添加fly方法则子类都具备该功能，但不是所有鸭子都能飞，比如橡皮鸭。如果让子类实现自己的fly方法又会出现上面的问题。
 
### 解决方案
行为的实现不应该跟鸭子类型耦合。因为不同的鸭子对应的行为可能相同，可能不同，如果完全相同，则这些行为可以放在父类中，比如swim，如果完全不同，倒是可以放在子类实现。既然不满足这两点，行为就应当拆出来，做成可以灵活更改的变量。

### UML类图
![design](https://www.plantuml.com/plantuml/png/0/ZPBBgi8m48RtynHT7XVr2KH4HCx6Wzf5bpBKU0ochUQY5FNT3TDRHXUs44Z--VwPIQOpOr5Rbol851eJBLnsYkuG-M-fghauubMMEZhq-nq5DmwpKnU-XAST497SPFupekr2chF2gZcjXJwKEW_IDu54s4-neLczdd-Ndz3SOojm_-jNS9--YoUGF7d4CGDRymgLGltT6t0080lhUyLCi9Uh_VxzUcySJqEahUgbLDfZAej_ysOB4quUup5Z0_zVcM8rm4iSQvdFiyv0K-LjCwmjJEA7XsUaoYIrhkJY10worMGt76V5u8eT33DHxFrl_W80 "design")


### 代码
```go
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

```

### 定义
strage pattern, 策略模式: 定义了一组算法，使其可以互相替换，将行为与类型解耦。

### 扩展
相比于继承，组合能带来更灵活的复用能力，可在运行时动态地改变行为，has-a关系要优于is-a关系。

### 原则
- 针对接口编程，不针对实现编程
- 多用组合，少用继承