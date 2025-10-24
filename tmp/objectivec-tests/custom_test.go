// Custom tests for ObjectiveC framework APIs - macOS compatible tests
// Tests basic object creation and non-accessibility features

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

	// Test basic object methods that work on macOS
	obj2 := objectivec.NewObject()
	if obj.IsEqualTo(obj2) {
		t.Log("Objects are equal (same singleton)")
	} else {
		t.Log("Objects are different instances")
	}
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
	// This test would normally test cross-framework type mapping
	// but many cross-framework methods are iOS-specific
	t.Skip("Cross-framework tests moved to iOS-specific test file")
}

// TestInterfaceCompliance verifies that objects implement IObject
func TestInterfaceCompliance(t *testing.T) {
	obj := objectivec.NewObject()

	// Verify object implements IObject
	var iobj objectivec.IObject = obj
	if iobj == nil {
		t.Fatal("Object does not implement IObject")
	}

	// Test that we can assign to interface
	t.Log("Interface compliance verified")
}

// TestNilHandling tests that nil/zero values are handled correctly
func TestNilHandling(t *testing.T) {
	obj := objectivec.NewObject()

	// Test with nil argument
	result := obj.IsEqualTo(objectivec.Object{})
	t.Logf("IsEqualTo with zero object: %v", result)
}

// TestMethodChaining tests that returned objects can be used for chaining
func TestMethodChaining(t *testing.T) {
	// Chain initialization calls
	obj := objectivec.ObjectClass.Alloc().Init()
	if obj.ID == 0 {
		t.Fatal("Method chaining failed")
	}
	t.Log("Method chaining works correctly")
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
