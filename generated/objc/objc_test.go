package objc

import (
	"testing"

	purego "github.com/ebitengine/purego/objc"
)

// BenchmarkSelCached benchmarks the cached Sel() function
func BenchmarkSelCached(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sel("alloc")
	}
}

// BenchmarkRegisterNameDirect benchmarks direct objc.RegisterName calls
func BenchmarkRegisterNameDirect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = purego.RegisterName("alloc")
	}
}

// BenchmarkSelCachedParallel benchmarks the cached Sel() function with parallel access
func BenchmarkSelCachedParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Sel("alloc")
		}
	})
}

// BenchmarkRegisterNameDirectParallel benchmarks direct objc.RegisterName with parallel access
func BenchmarkRegisterNameDirectParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = purego.RegisterName("alloc")
		}
	})
}

// BenchmarkSelCachedMultipleSelectors benchmarks cache with different selectors
func BenchmarkSelCachedMultipleSelectors(b *testing.B) {
	selectors := []string{"alloc", "init", "new", "autorelease", "length", "bytes"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sel(selectors[i%len(selectors)])
	}
}

// BenchmarkRegisterNameMultipleSelectors benchmarks RegisterName with different selectors
func BenchmarkRegisterNameMultipleSelectors(b *testing.B) {
	selectors := []string{"alloc", "init", "new", "autorelease", "length", "bytes"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = purego.RegisterName(selectors[i%len(selectors)])
	}
}
