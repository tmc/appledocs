// Custom tests for ObjectiveC framework APIs
// Tests various patterns including array returns, object returns, and property accessors

package objectivec_test

import (
	"testing"

	"github.com/tmc/appledocs/generated/objectivec"
)

// TestObjectCreation tests basic NSObject creation and method calls
func TestObjectCreation(t *testing.T) {
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("Failed to create NSObject instance")
	}

	// Test that we can call methods
	count := obj.AccessibilityElementCount()
	t.Logf("AccessibilityElementCount: %d", count)

	// Test boolean method
	result := obj.AccessibilityActivate()
	t.Logf("AccessibilityActivate result: %v", result)
}

// TestArrayReturnTypes tests methods that return arrays of objects
func TestArrayReturnTypes(t *testing.T) {
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

// TestObjectReturnTypes tests methods that return single objects
func TestObjectReturnTypes(t *testing.T) {
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

// TestFunctionReturnTypes tests methods that return function/block types
func TestFunctionReturnTypes(t *testing.T) {
	obj := objectivec.NewObject()

	// Test method returning function type
	// This tests that function types are preserved
	block := obj.AccessibilityExpandedStatusBlock()
	// Block will be nil for base NSObject, which is expected
	t.Logf("AccessibilityExpandedStatusBlock type: %T", block)
}

// TestPropertyAccessors tests property getters and setters
func TestPropertyAccessors(t *testing.T) {
	obj := objectivec.NewObject()

	// Test array property getter
	labels := obj.AccessibilityAttributedUserInputLabels()
	t.Logf("AccessibilityAttributedUserInputLabels length: %d", len(labels))

	// Test array property setter
	newLabels := make([]objectivec.IObject, 0)
	obj.SetAccessibilityAttributedUserInputLabels(newLabels)
	t.Log("Array property setter completed")
}

// TestClassMethods tests class-level methods
func TestClassMethods(t *testing.T) {
	// Test class method that returns an object
	cls := objectivec.ObjectClass

	// Test allocation
	obj := cls.Alloc()
	if obj.ID == 0 {
		t.Fatal("Alloc returned invalid object")
	}

	// Test initialization
	initialized := obj.Init()
	if initialized.ID == 0 {
		t.Fatal("Init returned invalid object")
	}
}

// TestCrossFrameworkTypes tests handling of types from other frameworks
func TestCrossFrameworkTypes(t *testing.T) {
	obj := objectivec.NewObject()

	// These methods return types from Foundation or other frameworks
	// In ObjectiveC framework context, they should map to IObject

	// Test method that would return foundation.Number in other contexts
	indices := obj.IndicesOfObjectsByEvaluatingObjectSpecifier(obj)
	if indices == nil {
		t.Log("IndicesOfObjectsByEvaluatingObjectSpecifier returned nil (expected)")
	} else {
		// Verify array element type
		for i, idx := range indices {
			if idx != nil {
				// Just verify we can call methods on array elements
				_ = idx.AccessibilityElementCount()
				t.Logf("Array element %d is valid IObject", i)
			}
		}
	}
}

// TestInterfaceCompliance verifies that returned objects implement IObject
func TestInterfaceCompliance(t *testing.T) {
	obj := objectivec.NewObject()

	// Get various objects and verify they implement IObject
	var iobj objectivec.IObject

	iobj = obj.AccessibilityElementAtIndex(0)
	if iobj == nil {
		t.Log("AccessibilityElementAtIndex returned nil (expected for base NSObject)")
	} else {
		// Test that we can call interface methods on returned objects
		count := iobj.AccessibilityElementCount()
		t.Logf("Object element count: %d", count)
	}
}

// TestArrayElementWrapping verifies array elements are properly wrapped
func TestArrayElementWrapping(t *testing.T) {
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

// TestNilHandling tests that nil/zero values are handled correctly
func TestNilHandling(t *testing.T) {
	obj := objectivec.NewObject()

	// Many accessors will return nil for base NSObject
	elem := obj.AccessibilityElementAtIndex(0)
	if elem != nil {
		t.Log("AccessibilityElementAtIndex is not nil (unusual but valid)")
	}

	// Arrays might be empty or nil
	labels := obj.AccessibilityAttributedUserInputLabels()
	if labels != nil && len(labels) > 0 {
		t.Log("AccessibilityAttributedUserInputLabels is not empty (unusual but valid)")
	}
}

// TestMethodChaining tests that returned objects can be used for chaining
func TestMethodChaining(t *testing.T) {
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

// BenchmarkObjectCreation benchmarks object creation performance
func BenchmarkObjectCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := objectivec.NewObject()
		if obj.ID == 0 {
			b.Fatal("Failed to create object")
		}
	}
}

// BenchmarkArrayReturn benchmarks array return type conversion
func BenchmarkArrayReturn(b *testing.B) {
	obj := objectivec.NewObject()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		labels := obj.AccessibilityAttributedUserInputLabels()
		_ = labels
	}
}

// BenchmarkObjectReturn benchmarks single object return type conversion
func BenchmarkObjectReturn(b *testing.B) {
	obj := objectivec.NewObject()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		elem := obj.AccessibilityElementAtIndex(0)
		_ = elem
	}
}
// TestClassGetName tests getting a class name using the Objective-C runtime
func TestClassGetName(t *testing.T) {
	// Create an NSObject instance
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("Failed to create NSObject instance")
	}

	// Get the class using the Objective-C runtime via reflection
	// We call the 'class' method on the object to get its class
	cls := obj.ID.Class()
	if cls == 0 {
		t.Fatal("Failed to get class from object")
	}

	t.Logf("Got class: %v", cls)
	
	// Note: class_getName is not exported from the objectivec package
	// In production code, you would typically use NSStringFromClass or 
	// access the class name through other means. This test demonstrates
	// that we can get the class successfully.
}
