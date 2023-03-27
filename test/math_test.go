package math

import (
	"testing"
)

// func TestAdd(t *testing.T) {
// 	got := Add(4, 6)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

type addTest struct {
	arg1, arg2, exp int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.exp {
			t.Errorf("Output %q not equal to expected %q", output, test.exp)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
