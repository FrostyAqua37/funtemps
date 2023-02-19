package main 

import (
	"fmt"
	"flag"

	"Github/funtemps/conv"
)

var fahrenheit float64
var celsius float64
var kelvin float64
var out string
var funfacts string

func init () {

flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i fahrenheit")
flag.Float64Var(&celsius, "C", 0.0, "temperatur i celsius")
flag.Float64Var(&kelvin, "K", 0.0, "temperatur i kelvin")
flag.StringVar(&out, "out", "C", "regner ut temperatur i C - celsius, F - fahrenheit, K - Kelvin")
flag.StringVar("funfacts", "funfacts", "sun", "\"fun-facts\" om sun - Sola, luna - Månen og terra - jorden")

}

func main() {

    flag.Parse()

	fmt.Println(fahrenheit, out, funfacts)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.Nflag()", flag.Nflag())

	f := "°Fahrenheit"
	c := "°Celsius"
	k := "Kelvin"

	//Omgjøring fra Fahrenheit Til Celsius.

	switch Float64Var { 
		case "F"
			switch StringVar {
                case "Celsius":
                    fmt.Printf("%.2f grader Fahrenheit = %.2f grader Celsius", Fahrenheit, Celsius)
                case "Kelvin":
					fmt.Printf("%.2f grader Fahrenheit = %.2f grader Kelvin", Fahrenheit, Kelvin)
				default:
				    fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
					return
            }
		case "C":
            switch StringVar {
				case "Fahrenheit":
				    fmt.Printf("%.2f grader Celsius = %.2f grader Fahrenheit", Celsius, Fahrenheit)
                case "Kelvin":
					fmt.Printf("%.2f grader Celsius = %.2f grader Kelvin", Celsius, Kelvin)
				default: 
					fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
					return
			}
		case "K":
            switch StringVar {	
				case "Fahrenheit":
				    fmt.Printf("%.2f grader Kelvin = %.2f grader Fahrenheit", Kelvin, fahrenheit)
                case "Celsius":
					fmt.Printlnfmt.Printf("%.2f grader Kelvin = %.2f grader Celsius", Kelvin, Celsius)
				default: 
					fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
				    return
            }
	}
}
