package fibonacci

import "testing"

func TestFiboRecursive(t *testing.T) {
	f := fiboRecursive(10)
	if f != 55 {
		t.Fatalf("Wrong fiboncci sequence, got %d, it should be 55.", f)
	}
}

func BenchmarkFiboRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboRecursive(10)
	}
}

func TestFiboClosure(t *testing.T) {
	f := fiboClosure()
	var out int
	for x := 0; x < 10; x++ {
		out = f()
	}
	if out != 55 {
		t.Fatalf("Wrong fiboncci sequence, got %d, it should be 55.", out)
	}
}

func BenchmarkFiboClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := fiboClosure()
		for x := 0; x < 10; x++ {
			f()
		}
	}
}

func TestFiboConcurrent(t *testing.T) {
	f := fiboConcurrent(10)
	if f != 55 {
		t.Fatalf("Wrong fiboncci sequence, got %d, it should be 55.", f)
	}
}

func BenchmarkFiboConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboConcurrent(10)
	}
}
