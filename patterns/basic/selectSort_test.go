package basic

import "testing"

func TestSelectSortExec(t *testing.T) {
	SelectSortExec()
}

func BenchmarkSelectSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		QuickSortExec()
	}
}
