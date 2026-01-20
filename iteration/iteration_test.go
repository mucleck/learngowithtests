package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeated := Repeat('a', 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	//everything here is ignored 
	for b.Loop() {
		Repeat('a', 5)
	}
	//or here
}

func ExampleRepeat() {
	characters := Repeat('a', 5)
	fmt.Println(characters)
	// Output: aaaaa
}

//To test and see allocs/op 
//go test -bench=. -benchmem

//add number of times it should run by parameter

//write ExampleRepeat

//search in strings for more useful methods