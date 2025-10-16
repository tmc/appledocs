package objc

import (
	"testing"

	"github.com/ebitengine/purego/objc"
)

func TestToNSString(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"simple string", "Hello"},
		{"string with spaces", "Hello, World!"},
		{"unicode", "Hello, 世界! 🌍"},
		{"multiline", "Line 1\nLine 2\nLine 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nsString := ToNSString(tt.input)
			if nsString == 0 {
				t.Fatal("ToNSString returned nil")
			}

			// Verify it's actually an NSString by calling length
			length := StringLength(nsString)
			if tt.input == "" && length != 0 {
				t.Errorf("expected length 0 for empty string, got %d", length)
			}
		})
	}
}

func TestToGoString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"simple string", "Hello", "Hello"},
		{"string with spaces", "Hello, World!", "Hello, World!"},
		{"unicode", "Hello, 世界!", "Hello, 世界!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert Go -> NSString -> Go
			nsString := ToNSString(tt.input)
			result := ToGoString(nsString)

			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestToGoStringNil(t *testing.T) {
	result := ToGoString(0)
	if result != "" {
		t.Errorf("expected empty string for nil NSString, got %q", result)
	}
}

func TestStringLength(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedLength uint64
	}{
		{"empty", "", 0},
		{"5 chars", "Hello", 5},
		{"with spaces", "Hello World", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nsString := ToNSString(tt.input)
			length := StringLength(nsString)

			if length != tt.expectedLength {
				t.Errorf("expected length %d, got %d", tt.expectedLength, length)
			}
		})
	}
}

func TestStringLengthNil(t *testing.T) {
	length := StringLength(0)
	if length != 0 {
		t.Errorf("expected length 0 for nil NSString, got %d", length)
	}
}

// BenchmarkToNSString benchmarks string to NSString conversion
func BenchmarkToNSString(b *testing.B) {
	str := "Hello, World!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToNSString(str)
	}
}

// BenchmarkToGoString benchmarks NSString to Go string conversion
func BenchmarkToGoString(b *testing.B) {
	nsString := ToNSString("Hello, World!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToGoString(nsString)
	}
}

// BenchmarkRoundTrip benchmarks full conversion cycle
func BenchmarkRoundTrip(b *testing.B) {
	str := "Hello, World!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nsString := ToNSString(str)
		_ = ToGoString(nsString)
	}
}

// TestUppercaseConversion demonstrates using objc methods with our conversions
func TestUppercaseConversion(t *testing.T) {
	input := "hello world"
	expected := "HELLO WORLD"

	// Convert to NSString
	nsString := ToNSString(input)

	// Call uppercaseString method
	sel_uppercaseString := objc.RegisterName("uppercaseString")
	upperNSString := nsString.Send(sel_uppercaseString)

	// Convert back to Go string
	result := ToGoString(upperNSString)

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
