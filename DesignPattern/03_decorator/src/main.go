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
