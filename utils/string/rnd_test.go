package uString

import "testing"

func BenchmarkRnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rnd(9)
	}
	// 242ns
}

func BenchmarkRnd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rnd2(9)
	}
	//400
}
func BenchmarkRnd3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rnd3(9)
	}
	//203
}

func BenchmarkRnd4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rnd4(9)
	}
	//207
}
