package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {

	Test(ShellSort, t)
}

func BenchmarkShellSort(b *testing.B) {

	Benchmark(ShellSort, b)
}
