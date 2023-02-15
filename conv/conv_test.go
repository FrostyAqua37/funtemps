package conv

import (
	"reflect"
	"testing"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

<<<<<<< HEAD
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64;
		want  float64;
	}
	tests := []test{
		{input_, want:__},
	}

	for _, tc :=range tests {
		got := KelvinToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestCelsiusToFahrenheit
func TestCelsiusToKelvin
func TestKelvinToFahrenheit
func TestFahrenheitToKelvin


// De andre testfunksjonene implementeres her
// ...
=======
// De andre testfunksjonene implementeres her
// ...
>>>>>>> 5d7ecb0e3ba2673b281dfa6f4f46732e43eca334
