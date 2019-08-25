package calc

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 1)
	if result != 2 {
		t.Errorf("Expecting result %v, got %v", 2, result)
	}

}

// func TestAddVaradic (t *testing.T) {
// 	result := struct[] {
// 		nums ..int
// 		res
// 	}
// 	if result != 2 {
// 		t.Errorf("Expecting result %v, got %v", 2, result)
// 	}

// }