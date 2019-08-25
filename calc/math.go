package calc

// Add sums two values returning result
// func Add(a int, b int) int {
// 	res := a + b
// 	return res
// }
func Add(numbers ...int) int {
	sum :=0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
