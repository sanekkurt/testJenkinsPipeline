package funcs

import "testing"

func TestSum(t *testing.T) {

	got := Sum(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
