### 背景
1. 设计一个咖啡饮料订单系统，能够获取每一种饮料的价格和描述
2. 饮料由咖啡和调料组成，每种咖啡可以搭配多种调料
2. 现有的设计：每种咖啡和调料的组合都生成一个类，单独生成价格和描述

### 问题
1. 高耦合，咖啡和调料静态绑定后直接导致类爆炸
2. 类爆炸直接导致一系列的开发维护问题

### 解决方案
将咖啡和调料从静态绑定换成动态组合

### UML类图
![design](https://www.plantuml.com/plantuml/png/0/SoWkIImgAStDuShCAqajIajCJbLmIYrBBKfCJrMevb800lVK4fVKukIYp8AIpFmy3KsOe7D-SGcGLyl5bPoJM5oiu9oVbfQQQaYiBhWjDZLwUZ3Dg6OPKMAM4SmibzJa_AB4almYtyoSRAY9S0r9v5Y0OYA6DS69EINKjGPa8XUNGsfU2Z2Y0000 "design")


### 代码
```golang
package main

import "fmt"

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

type Coffee struct {
	desc string
	cost float64
}

func NewCoffee(desc string, cost float64) *Coffee {
	return &Coffee{desc, cost}
}

func (c *Coffee) GetDescription() string {
	return c.desc
}

func (c *Coffee) GetCost() float64 {
	return c.cost
}

type Decorator struct {
	b Beverage
	Coffee
}

func NewDecorator(desc string, cost float64, b Beverage) *Decorator {
	return &Decorator{
		b:      b,
		Coffee: Coffee{desc: desc, cost: cost},
	}
}

func (d *Decorator) GetDescription() string {
	return d.b.GetDescription() + ", " + d.desc
}

func (d *Decorator) GetCost() float64 {
	return d.b.GetCost() + d.cost
}

func main() {
	b1 := NewDecorator("milk", .1, NewDecorator("salt", .2, NewCoffee("coffee1", 1000)))
	fmt.Printf("b1:%#v cost: $%#v\n", b1.GetDescription(), b1.GetCost())

	var b2 Beverage = NewCoffee("coffee2", 2000)
	b2 = NewDecorator("milk", .2, b2)
	b2 = NewDecorator("milk", .2, b2)
	b2 = NewDecorator("salt", 1.0, b2)
	fmt.Printf("b2:%#v cost: $%#v\n", b2.GetDescription(), b2.GetCost())
}
```

### 定义
Decorator, 装饰者模式：动态地将责任附加到对象上

### 扩展
decorator模式不使用继承来扩展对象功能，通过不停地装饰来动态增加责任

### 原则
- 开闭原则
- 针对接口编程，不针对实现编程