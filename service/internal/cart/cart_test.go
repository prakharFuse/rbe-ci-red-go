package cart

import "testing"

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3}); got != 6 {
		t.Fatalf("Total([1 2 3]) = %d, want 6", got)
	}
}
