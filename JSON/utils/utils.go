package utils

// SumaPares suma los números pares hasta el número dado
func SumaPares(n int) int {
	sum := 0
	for i := 0; i <= n; i += 2 {
		sum += i
	}
	return sum
}

// SumaImpares suma los números impares has
func SumaImpares(n int) int {
	sum := 0
	for i := 1; i <= n; i += 2 {
		sum += i
	}
	return sum
}