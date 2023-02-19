package conv

import (
	"reflect"
	"testing"
)


func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		  input float64
		  want 	float64
	}

	tests := []test{
		  {input: 150, want: 65.56},
	}

	for _, tc := range tests {
  		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    		t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		  input float64
          want 	float64
    }
	tests := []test{
          {input: 150, want: 338.71},
	}
	for _, tc := range tests {
  		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    		t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  		}	
	}
}

 func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 150, want: -123.15},
	}

		for _, tc := range tests {
  		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    	t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
        want  float64
    }
	tests := []test{
        {input: 150, want: -189.67},
	}
	for _, tc := range tests {
  		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    	t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 150, want: 302},
	}

	for _, tc := range tests {
 		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    	t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
  }
}
}

func TestCelsiusToKelvin(t *testing.T) {

	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 150, want: 423.15},
	}

	for _, tc := range tests {
  		got := FarhenheitToCelsius(tc.input)
  		if !withinTolerance(tc.want, got, 1e-12) {
    		t.Errorf("expected: %.18f, got: %.18f", tc.want, got)
 		}	
	}
}

	



