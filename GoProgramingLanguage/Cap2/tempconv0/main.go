package main

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsolutZeroC Celcius = -273.15
	FreezingC    Celcius = 0
	BoilingC     Celcius = 100
)

func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func (c Celcius) String() string { return fmt.Sprintf("%gºC", c) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//fmt.Printf("%g\n", boilingF-FreezingC)

	var c Celcius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f)
	fmt.Println(c == Celcius(f))

	c = FToC(212.0)

	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)

	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
