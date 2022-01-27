package tempconv

import "fmt"

type Celsius float64
type Kelvin float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	zeroK         Kelvin  = 0
)

func CelsiusToFahrenheit(c Celsius) (f Fahrenheit, inputC Celsius) {
	return Fahrenheit(c*9/5 + 32), c
}

func FahrenheitToCelsius(f Fahrenheit) (c Celsius, inputF Fahrenheit) {
	return Celsius((f - 32) * 5 / 9), f
}

func CelsiusToKelvin(c Celsius) (k Kelvin, inputC Celsius) {
	return Kelvin(c + 273.15), c
}

func FahrenheitToKelvin(f Fahrenheit) (k Kelvin, inputF Fahrenheit) {
	return Kelvin((f + 459.67) / 1.8), f
}

func KelvinToCelsius(k Kelvin) (c Celsius, inputK Kelvin) {
	return Celsius(k - 273.15), k
}

func KelvinToFahrenheit(k Kelvin) (f Fahrenheit, inputK Kelvin) {
	return Fahrenheit(k*1.8 - 459.67), k
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f°K", k)
}
