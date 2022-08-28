package calc

import (
	"errors"
)

func Add(x, y float64) (float64, error) {
	return x + y, nil
}

func Sub(x, y float64) (float64, error) {
	return x - y, nil
}

func Mult(x, y float64) (float64, error) {
	return x * y, nil
}
func Div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New(`invalid value.dividing number can not be "0"`)
	}
	return x / y, nil
}
