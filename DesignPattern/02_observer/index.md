### 背景
1. 设计一个天气状态服务器，实时监测温度、湿度和气压，并将最新状态同步给订阅者。
2. 旧设计方案：将每个订阅者的更新代码写死到服务端代码里。

### 问题
1. 高耦合，服务端代码需要知道订阅者的实现
2. 违反了开闭原则，新加订阅者需要修改服务端代码

### 解决方案
约定发布者跟订阅者之间的接口，抽象出注册、注销、通知这三个方法，发布者暴露出这三个方法给订阅者使用即可

### UML类图
![UML类图](https://github.com/changmu/NoteBook/blob/master/DesignPattern/02_observer/src/uml.png?raw=true)


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