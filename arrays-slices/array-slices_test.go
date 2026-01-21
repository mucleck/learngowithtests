package arraysslices

import "testing"

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