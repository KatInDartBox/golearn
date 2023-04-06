package uJson

import "testing"

func BenchmarkFilterUser1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FilterUser1()
	}
	//0.2555 ns/op
}

func BenchmarkFilterUser2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FilterUser2()
	}
	//1119 ns/op
}
