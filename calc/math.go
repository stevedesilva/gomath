package calc

import "errors"

// Add sums two values returning result
//
// v1:
// func Add(a int, b int) int {
// 	res := a + b
// 	return res
// }
//
// v2
// func Add(numbers ...int) int {
// 	sum :=0
//
// 	for _, num := range numbers {
// 		sum += num
// 	}
// 	return sum
// }
//
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Provide more that 2 numbers")
	}

	for _, num := range numbers {
		sum += num
	}

	return sum, nil

}
