package lib

func Addition(a int, b int) float64 {
	return float64(a + b)
}

func Subtraction(a int, b int) float64 {
	return float64(a - b)
}

func Multiplication(a int, b int) float64 {
	return float64(a * b)
}

func Division(a int, b int) float64 {
	if b == 0 {
		return float64(0)
	}
	return float64(a / b)
}
