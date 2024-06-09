package main

import "fmt"

func main() {

	var x AmericanVelocity = 120.4
	fmt.Println(x.ToEuropean())
	fmt.Println(x)
	ev := x.ToEuropean()
	fmt.Println(ev.ToAmerican())

}

type AmericanVelocity float64
type EuropeanVelocity float64

const coef = 5793.624

func (v AmericanVelocity) ToEuropean() EuropeanVelocity {
	ev := EuropeanVelocity(v * coef)
	return ev
}

func (v AmericanVelocity) String() string {
	return fmt.Sprintf("%.1f m/sec", v)
}

func (v EuropeanVelocity) String() string {
	return fmt.Sprintf("%.1f km/hour", v)
}

func (v EuropeanVelocity) ToAmerican() AmericanVelocity {
	av := AmericanVelocity(v / coef)
	return av
}
