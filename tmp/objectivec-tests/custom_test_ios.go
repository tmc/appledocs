//go:build ios || tvos || visionos

// Custom tests for ObjectiveC framework APIs - iOS/tvOS/visionOS specific tests
// Tests accessibility features that are only available on iOS platforms

package objectivec_test

import (
	"testing"

	"github.com/tmc/appledocs/generated/objectivec"
)

// TestAccessibilityMethods tests iOS-specific accessibility methods on NSObject
func TestAccessibilityMethods(t *testing.T) {
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("Failed to create NSObject instance")
	}

	// Test accessibility methods
	count := obj.AccessibilityElementCount()
	t.Logf("AccessibilityElementCount: %d", count)

	// Test boolean method
	result := obj.AccessibilityActivate()
	t.Logf("AccessibilityActivate result: %v", result)
}

// TestAccessibilityArrayReturnTypes tests accessibility methods that return arrays
func TestAccessibilityArrayReturnTypes(t *testing.T) {
	obj := objectivec.NewObject()

	// Test property getter returning array
	// This tests the []objc.ID -> []IObject conversion
	labels := obj.AccessibilityAttributedUserInputLabels()
	if labels == nil {
		t.Log("AccessibilityAttributedUserInputLabels returned nil (expected for base NSObject)")
	}

	// Verify we can iterate over the result
	for i, label := range labels {
		if label == nil {
			t.Errorf("Array element %d is nil", i)
		}
	}
}

// TestAccessibilityObjectReturnTypes tests accessibility methods that return single objects
func TestAccessibilityObjectReturnTypes(t *testing.T) {
	obj := objectivec.NewObject()

	// Test method returning IObject
	// This tests the objc.ID -> IObject wrapping
	elem := obj.AccessibilityElementAtIndex(0)
	if elem == nil {
		t.Log("AccessibilityElementAtIndex returned nil (expected for base NSObject)")
	} else {
		// Test that returned object is valid - just verify we can call methods on it
		_ = elem.AccessibilityElementCount()
		t.Log("AccessibilityElementAtIndex returned valid IObject")
	}
}

// TestAccessibilityFunctionReturnTypes tests methods that return function/block types
func TestAccessibilityFunctionReturnTypes(t *testing.T) {
	obj := objectivec.NewObject()

	// Test method returning function type
	// This tests that function types are preserved
	block := obj.AccessibilityExpandedStatusBlock()
	// Block will be nil for base NSObject, which is expected
	t.Logf("AccessibilityExpandedStatusBlock type: %T", block)
}

// TestAccessibilityPropertyAccessors tests accessibility property getters and setters
func TestAccessibilityPropertyAccessors(t *testing.T) {
	obj := objectivec.NewObject()

	// Test array property getter
	labels := obj.AccessibilityAttributedUserInputLabels()
	t.Logf("AccessibilityAttributedUserInputLabels length: %d", len(labels))

	// Test array property setter
	newLabels := make([]objectivec.IObject, 0)
	obj.SetAccessibilityAttributedUserInputLabels(newLabels)
	t.Log("Array property setter completed")
}

// TestAccessibilityArrayElementWrapping verifies array elements are properly wrapped
func TestAccessibilityArrayElementWrapping(t *testing.T) {
	obj := objectivec.NewObject()

	// Get an array of objects
	labels := obj.AccessibilityAttributedUserInputLabels()

	// Each element should implement IObject interface
	for i, label := range labels {
		if label == nil {
			continue
		}

		// Verify we can call interface methods
		_ = label.AccessibilityElementCount()
		_ = label.AccessibilityActivate()

		t.Logf("Array element %d implements IObject correctly", i)
	}
}

// TestAccessibilityMethodChaining tests that returned objects can be used for chaining
func TestAccessibilityMethodChaining(t *testing.T) {
	obj := objectivec.NewObject()

	// Get an element and call methods on it
	elem := obj.AccessibilityElementAtIndex(0)
	if elem != nil {
		count := elem.AccessibilityElementCount()
		t.Logf("Chained call result: %d", count)
	} else {
		t.Log("Element is nil (expected for base NSObject), method chaining would work if not nil")
	}
}

// BenchmarkAccessibilityArrayReturn benchmarks array return type conversion
func BenchmarkAccessibilityArrayReturn(b *testing.B) {
	obj := objectivec.NewObject()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		labels := obj.AccessibilityAttributedUserInputLabels()
		_ = labels
	}
}

// BenchmarkAccessibilityObjectReturn benchmarks single object return type conversion
func BenchmarkAccessibilityObjectReturn(b *testing.B) {
	obj := objectivec.NewObject()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		elem := obj.AccessibilityElementAtIndex(0)
		_ = elem
	}
}
