package cal

import "testing"

func TestSumOfFirstThree(t *testing.T) {
	given := 3
	want := 6

	get := sumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but get %d\n", given, want, get)
	}
}