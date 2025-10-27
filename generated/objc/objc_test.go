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

// TestString tests the String helper function
func TestString(t *testing.T) {
	// Test that String conversion works without crashing
	str := String("Hello, World!")
	if str == 0 {
		t.Error("Expected non-zero NSString ID")
	}

	// Test empty string
	emptyStr := String("")
	if emptyStr == 0 {
		t.Error("Expected non-zero NSString ID for empty string")
	}
}

// BenchmarkString benchmarks the String conversion function
func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = String("benchmark string")
	}
}

