package basic

import "testing"

//go test -v -run TestBinSearchExec
func TestBinSearchExec(t *testing.T) {
	BinSearchExec()
}

func BenchmarkBinSearchExec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinSearchExec2()
	}
}
