package util

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x
	}

	tmp := Pow(x, y/2)
	if y%2 == 0 {
		return tmp * tmp
	}
	return x * tmp * tmp
}
