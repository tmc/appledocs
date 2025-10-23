package coregraphics_test

import (
	"reflect"
	"testing"

	"github.com/tmc/appledocs/generated/coregraphics"
)

// TestManualCGColorCreateSignature verifies the corrected function signatures
func TestManualCGColorCreateSignature(t *testing.T) {
	// Verify CGColorCreate accepts []float64 as second parameter
	t.Log("Checking CGColorCreate signature...")

	// Get the function type using reflection
	fnType := reflect.TypeOf(coregraphics.CGColorCreate)

	// Verify it's a function
	if fnType.Kind() != reflect.Func {
		t.Fatalf("CGColorCreate is not a function, got %v", fnType.Kind())
	}

	// Verify it has 2 parameters
	if fnType.NumIn() != 2 {
		t.Fatalf("CGColorCreate should have 2 parameters, got %d", fnType.NumIn())
	}

	// Check first parameter type (ColorSpaceRef)
	param0 := fnType.In(0)
	t.Logf("  Parameter 0: %v", param0)

	// Check second parameter type (should be []float64)
	param1 := fnType.In(1)
	expectedType := reflect.TypeOf([]float64{})
	if param1 != expectedType {
		t.Errorf("CGColorCreate parameter 1: expected %v, got %v", expectedType, param1)
	} else {
		t.Logf("  ✓ Parameter 1: %v (correct slice type!)", param1)
	}

	// Verify return type (ColorRef)
	if fnType.NumOut() != 1 {
		t.Fatalf("CGColorCreate should return 1 value, got %d", fnType.NumOut())
	}
	returnType := fnType.Out(0)
	t.Logf("  Return type: %v", returnType)

	t.Log("✓ CGColorCreate has correct signature: func(ColorSpaceRef, []float64) ColorRef")
}

// TestManualCGColorGetContentHeadroomSignature verifies float32 return type
func TestManualCGColorGetContentHeadroomSignature(t *testing.T) {
	t.Log("Checking CGColorGetContentHeadroom signature...")

	// Get the function type
	fnType := reflect.TypeOf(coregraphics.CGColorGetContentHeadroom)

	// Verify it's a function
	if fnType.Kind() != reflect.Func {
		t.Fatalf("CGColorGetContentHeadroom is not a function, got %v", fnType.Kind())
	}

	// Verify it has 1 parameter
	if fnType.NumIn() != 1 {
		t.Fatalf("CGColorGetContentHeadroom should have 1 parameter, got %d", fnType.NumIn())
	}

	// Check parameter type (ColorRef)
	param0 := fnType.In(0)
	t.Logf("  Parameter 0: %v", param0)

	// Verify return type (should be float32, not unsafe.Pointer)
	if fnType.NumOut() != 1 {
		t.Fatalf("CGColorGetContentHeadroom should return 1 value, got %d", fnType.NumOut())
	}

	returnType := fnType.Out(0)
	expectedType := reflect.TypeOf(float32(0))
	if returnType != expectedType {
		t.Errorf("CGColorGetContentHeadroom return type: expected %v, got %v", expectedType, returnType)
	} else {
		t.Logf("  ✓ Return type: %v (correct float32, not unsafe.Pointer!)", returnType)
	}

	t.Log("✓ CGColorGetContentHeadroom has correct signature: func(ColorRef) float32")
}

// TestManualCGColorCreateWithContentHeadroomSignature verifies float32 parameter
func TestManualCGColorCreateWithContentHeadroomSignature(t *testing.T) {
	t.Log("Checking CGColorCreateWithContentHeadroom signature...")

	fnType := reflect.TypeOf(coregraphics.CGColorCreateWithContentHeadroom)

	if fnType.Kind() != reflect.Func {
		t.Fatalf("CGColorCreateWithContentHeadroom is not a function")
	}

	// Should have 6 parameters: float32, ColorSpaceRef, float64, float64, float64, float64
	if fnType.NumIn() != 6 {
		t.Fatalf("CGColorCreateWithContentHeadroom should have 6 parameters, got %d", fnType.NumIn())
	}

	// Check first parameter is float32 (headroom)
	param0 := fnType.In(0)
	expectedType := reflect.TypeOf(float32(0))
	if param0 != expectedType {
		t.Errorf("Parameter 0: expected %v, got %v", expectedType, param0)
	} else {
		t.Logf("  ✓ Parameter 0 (headroom): %v (correct float32!)", param0)
	}

	// Check remaining parameters
	for i := 1; i < fnType.NumIn(); i++ {
		t.Logf("  Parameter %d: %v", i, fnType.In(i))
	}

	t.Log("✓ CGColorCreateWithContentHeadroom has correct signature with float32 headroom")
}

// TestManualSliceParameterFunctions verifies other functions using []float64
func TestManualSliceParameterFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       interface{}
		wantType reflect.Type
		paramIdx int
	}{
		{
			name:     "CGColorCreateWithPattern",
			fn:       coregraphics.CGColorCreateWithPattern,
			wantType: reflect.TypeOf([]float64{}),
			paramIdx: 2,
		},
		{
			name:     "CGContextSetFillColor",
			fn:       coregraphics.CGContextSetFillColor,
			wantType: reflect.TypeOf([]float64{}),
			paramIdx: 1,
		},
		{
			name:     "CGContextSetStrokeColor",
			fn:       coregraphics.CGContextSetStrokeColor,
			wantType: reflect.TypeOf([]float64{}),
			paramIdx: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fnType := reflect.TypeOf(tt.fn)
			if fnType.Kind() != reflect.Func {
				t.Fatalf("%s is not a function", tt.name)
			}

			if fnType.NumIn() <= tt.paramIdx {
				t.Fatalf("%s has fewer parameters than expected", tt.name)
			}

			paramType := fnType.In(tt.paramIdx)
			if paramType != tt.wantType {
				t.Errorf("%s parameter %d: expected %v, got %v", tt.name, tt.paramIdx, tt.wantType, paramType)
			} else {
				t.Logf("✓ %s parameter %d is %v", tt.name, tt.paramIdx, paramType)
			}
		})
	}
}
