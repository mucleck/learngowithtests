package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	
	// run go test -cover to see how much in percentage are you covering with your
	// tests, the goal is not to make for tests for the sake of it

	/*t.Run("aaray of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}) */

	t.Run("slice of n numbers", func(t *testing.T){
		numbers := []int{1,2,3,4,5,6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{2,2}, []int{2,3})
	want := []int{3, 4, 5}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func BenchmarkRepeatSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1,2}, []int{2,2}, []int{2,3})
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1,2}, []int{1,1,1}, []int{1,1}, []int{})
	want := []int{2, 2, 1, 0}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}