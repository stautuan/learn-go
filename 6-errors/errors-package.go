package main

import "errors"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// when returning a non-nil error in Go, it's conventional to return 0
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}