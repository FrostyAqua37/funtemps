package main 

import (
	"fmt"
	"flag"

	"github.com/FrostyAqua37/funtemps/conv"
)

var F float64
var C float64
var K float64
var out string
var funfacts string

func init () {

flag.Float64Var(&F, "F", 0.0, "temperatur i grader fahrenheit")
flag.Float64Var(&C, "C", 0.0, "temperatur i grader celsius")
flag.Float64Var(&K, "K", 0.0, "temperatur i kelvin")
flag.StringVar(&out, "out", "C", "regner ut temperatur i C - celsius, F - fahrenheit, K - Kelvin")
flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Sola, luna - Månen og terra - jorden")

}



func main() {
    flag.Parse()

	fmt.Println(F, out, funfacts)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.Nflag()", flag.NFlag())


	fmt.Println(F, C, K)
	var Res float64;

	//Omgjøring fra Fahrenheit til Celsius eller Kelvin.
	if(F !=  0) {
		switch out {
			case "C":
				if isFlagPassed("F") {
					Res = conv.FarhenheitToCelsius(C)
					fmt.Println("Fahrenheit til Celsius", Res)
				}

			case "K":
				if isFlagPassed("F") {
					Res := conv.FarhenheitToKelvin(K)
					fmt.Println("Fahrenheit til Kelvin", Res)
				}
		}
	} 

	//Omgjøring fra Celsius til Fahrenheit eller Kelvin. 
	if(C != 0) {
		switch out {
			case "F":
				if isFlagPassed("C") {
					Res = conv.CelsiusToFahrenheit(F)
					fmt.Println("Celsius til Fahrenheit", Res)
			}
			case "K":
				if isFlagPassed("C") {
					Res = conv.CelsiusToKelvin(K)
					fmt.Println("Celsius til Kelvin", Res)
			}
		}
	}
	//Omgjøring fra Kelvin til Celsius eller Fahrenheit. 
	if(K != 0) {
		switch out {
			case "C":
				if isFlagPassed("K") {
					Res = conv.KelvinToCelsius(C)
					fmt.Println("Kelvin til Celsius", Res)
			}
			case "F":
				if isFlagPassed("K") {
					Res = conv.KelvinToFahrenheit(F)
					fmt.Println("Kelvin til Fahrenheit", Res)
			}
		}
	}     
	
}

	func isFlagPassed(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		return found
	}