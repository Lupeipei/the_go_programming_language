package tempcov

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = -273.15
	BoilingC Celsius = -273.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5+32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f-32)*5/9) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// func main() {
// 	fmt.Printf("%g\n", BoilingC - FreezingC)
//
// 	boilingF := CToF(BoilingC)
//
// 	fmt.Printf("%g\n", boilingF - CToF(FreezingC))
//
// 	c := FToC(212.0)
// 	fmt.Println(c.String())
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%s\n", c)
// 	fmt.Println(c)
// 	fmt.Printf("%g\n", c)
// }
