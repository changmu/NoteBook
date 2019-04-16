### 背景
1. 设计一个咖啡饮料订单系统，能够获取每一种饮料的价格和描述
2. 饮料由咖啡和调料组成，每种咖啡可以搭配多种调料
2. 现有的设计：每种咖啡和调料的组合都生成一个类，单独生成价格和描述

### 问题
1. 高耦合，咖啡和调料静态绑定后直接导致类爆炸
2. 类爆炸直接导致一系列的开发维护问题

### 解决方案
将咖啡和调料同静态绑定换成动态组合

### UML类图



### 代码
```go
package main

import "fmt"

// 定义接口
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

type Observer interface {
	Update()
}

type DisplayElement interface {
	Display()
}

type WeatherData struct {
	observers []Observer

	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
	fmt.Printf("observer:%p registered.\n", o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	for idx, item := range w.observers {
		if item == o {
			w.observers = append(w.observers[:idx], w.observers[idx+1:]...) // attention
			fmt.Printf("observer:%p removed.\n", o)
			return
		}
	}
}

func (w *WeatherData) NotifyObserver() {
	for _, item := range w.observers {
		item.Update()
	}
}

func (w *WeatherData) GetTemperature() float64 {
	return w.temperature
}

func (w *WeatherData) GetHumidity() float64 {
	return w.humidity
}

func (w *WeatherData) GetPressure() float64 {
	return w.pressure
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObserver()
}

func (w *WeatherData) SetMeasurements(t, h, p float64) {
	w.temperature = t
	w.humidity = h
	w.pressure = p
	w.MeasurementsChanged()
}

// 订阅者
type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	// pressure    float64
	weatherData *WeatherData
}

func NewCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay {
	d := &CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(d)
	return d
}

func (d *CurrentConditionsDisplay) Update() {
	d.temperature = d.weatherData.GetTemperature()
	d.humidity = d.weatherData.GetHumidity()
	d.Display()
}

func (d *CurrentConditionsDisplay) Display() {
	fmt.Printf("current conditions: temperature:%v and humidity:%v\n", d.temperature, d.humidity)
}

func main() {
	weatherData := NewWeatherData()
	currentConditionsDisplay := NewCurrentConditionsDisplay(weatherData)
	defer weatherData.RemoveObserver(currentConditionsDisplay)
	currentConditionsDisplay2 := NewCurrentConditionsDisplay(weatherData)
	defer weatherData.RemoveObserver(currentConditionsDisplay2)

	weatherData.SetMeasurements(10, 20, 30)
	weatherData.SetMeasurements(100, 200, 300)
}
```

### 定义
Observer, 观察者模式：定义了对象之间的一对多依赖，当主题对象改变时，订阅者都将收到通知并更新。

### 扩展
订阅者模式可以说是一种单向依赖，订阅者依赖于发布者的状态获取接口，算是依赖倒置的应用，变化点在于订阅者

### 原则
- 为交互对象之间的松耦设计合而努力
- 开闭原则
- 针对接口编程，不针对实现编程