package aoc

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, others ...int) int {
	result := a * b / GCD(a, b)
	for i := range len(others) {
		result = LCM(result, others[i])
	}
	return result
}
