package array

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d, given %d", got, want, numbers)
	}
}

func Sum(numbers [5]int) int {
	Sum := 0
	for _, number := range numbers {
		Sum += number
	}
	return Sum

}
