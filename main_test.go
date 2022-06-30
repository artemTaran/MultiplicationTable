package main

import "testing"

//BenchmarkTable-8           13108             88946 ns/op
func BenchmarkTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable(100)
	}
}

//BenchmarkTable2-8          13122             91535 ns/op
func BenchmarkTable2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable2(100)
	}
}

//BenchmarkTable3-8          13681             88975 ns/op
//BenchmarkTable3-8          13191             91248 ns/op
//BenchmarkTable3-8          13348             89158 ns/op
func BenchmarkTable3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable3(100)
	}
}

//BenchmarkTable4-8          13063             90085 ns/op
//BenchmarkTable4-8          12892             89099 ns/op
//BenchmarkTable4-8          12946             87818 ns/op
func BenchmarkTable4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable4(100)
	}
}
