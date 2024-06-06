// Section 4: Testing

// Unit Testing for SumEven function
package section4

import (
	"testing"
	"go-test-exercise/section1"
)

func TestSumEven(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 12},
		{[]int{1, 3, 5}, 0},
		{[]int{2, 4, 6, 8}, 20},
		{[]int{}, 0},
	}

	for _, test := range tests {
		got := section1.SumEven(test.input)
		if got != test.want {
			t.Errorf("SumEven(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
