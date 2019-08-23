package calc

import "testing"

func TestMath(t *testing.T) {
	result := Add(1, 1)
	if result != 2 {
		t.Errorf("Expecting result %v, got %v", 2, result)
	}

}
