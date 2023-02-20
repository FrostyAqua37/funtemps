package main 

import (
	"fmt"
	"flag"

	"Github/funtemps/conv"
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

	fmt.Println(fahrenheit, out, funfacts)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.Nflag()", flag.Nflag())

	f := "°Fahrenheit"
	c := "°Celsius"
	k := "Kelvin"

	//Omgjøring fra Fahrenheit til Celsius eller Kelvin.

	switch Float64Var { 
	case "F":
			switch StringVar {
                case "Celsius":
                    fmt.Printf("%.2f grader Fahrenheit = %.2f grader Celsius", F, C)
                case "Kelvin":
					fmt.Printf("%.2f grader Fahrenheit = %.2f grader Kelvin", F, K)
				default:
				    fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
					return
            	}

		//Omgjøring fra Celsius til Fahrenheit eller Kelvin 
		case "C":
            switch StringVar {
				case "Fahrenheit":
				    fmt.Printf("%.2f grader Celsius = %.2f grader Fahrenheit", C, F)
                case "Kelvin":
					fmt.Printf("%.2f grader Celsius = %.2f grader Kelvin", C, K)
				default: 
					fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
					return
				}

		//Omgjøring fra Kelvin til Fahrenheit eller Kelvin
		case "K":
            switch StringVar {	
				case "Fahrenheit":
				    fmt.Printf("%.2f grader Kelvin = %.2f grader Fahrenheit", K, F)
                case "Celsius":
					fmt.Printf("%.2f grader Kelvin = %.2f grader Celsius", K, C)
				default: 
					fmt.Println("Error: Invalid output temperature type: %s\n", StringVar)
				    return
            	}
	}
}
