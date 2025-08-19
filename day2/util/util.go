package util

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) float64 {
	if b == 0 {
		return 0 // Avoid division by zero
	}
	return float64(a) / float64(b)
}
