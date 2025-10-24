// Comprehensive tests for NSObject basic functionality
// These tests verify actual NSObject methods that work at runtime.

package objectivec_test

import (
	"testing"

	"github.com/tmc/appledocs/generated/objectivec"
)

// TestObjectCreation tests basic NSObject creation and initialization
func TestObjectCreation(t *testing.T) {
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("Failed to create NSObject instance")
	}

	// Test Init method
	obj2 := obj.Init()
	if obj2.ID == 0 {
		t.Fatal("Init() returned invalid object")
	}
	t.Logf("Init() returned object: %v", obj2.ID)

	// Test IsEqualTo method
	equal := obj.IsEqualTo(obj2)
	t.Logf("IsEqualTo(self): %v", equal)
	if !equal {
		t.Error("Object should be equal to itself")
	}

	// Test Autorelease method
	obj3 := obj.Autorelease()
	if obj3.ID == 0 {
		t.Fatal("Autorelease() returned invalid object")
	}
	t.Logf("Autorelease() returned object: %v", obj3.ID)
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
	t.Logf("Alloc() created object: %v", obj.ID)

	// Test initialization
	initialized := obj.Init()
	if initialized.ID == 0 {
		t.Fatal("Init returned invalid object")
	}
	t.Logf("Init() returned object: %v", initialized.ID)
}

// TestNewObject tests the convenience constructor
func TestNewObject(t *testing.T) {
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("NewObject() failed")
	}
	t.Logf("NewObject() created: %v", obj.ID)
}

// TestObjectEquality tests object equality checking
func TestObjectEquality(t *testing.T) {
	obj1 := objectivec.NewObject()
	obj2 := objectivec.NewObject()

	// Test equality with self
	if !obj1.IsEqualTo(obj1) {
		t.Error("Object should be equal to itself")
	}

	// Test inequality with different object
	// Note: We can't assume obj1 != obj2 since they're both empty NSObjects
	// and Objective-C may return the same instance
	t.Logf("obj1.IsEqualTo(obj2): %v", obj1.IsEqualTo(obj2))
}

// TestObjectLifecycle tests object lifecycle methods
func TestObjectLifecycle(t *testing.T) {
	// Create object
	obj := objectivec.ObjectClass.Alloc()
	t.Logf("Allocated object: %v", obj.ID)

	// Initialize
	obj = obj.Init()
	t.Logf("Initialized object: %v", obj.ID)

	// Autorelease (adds to autorelease pool)
	obj = obj.Autorelease()
	t.Logf("Autoreleased object: %v", obj.ID)

	// Object should still be valid here since we're in an autorelease pool
	if obj.ID == 0 {
		t.Fatal("Object became invalid after autorelease")
	}
}

// TestMultipleObjects tests creating multiple objects
func TestMultipleObjects(t *testing.T) {
	objects := make([]objectivec.Object, 10)

	for i := 0; i < 10; i++ {
		objects[i] = objectivec.NewObject()
		if objects[i].ID == 0 {
			t.Fatalf("Failed to create object %d", i)
		}
	}

	t.Logf("Successfully created %d objects", len(objects))

	// Verify all objects are valid
	for i, obj := range objects {
		if obj.ID == 0 {
			t.Errorf("Object %d is invalid", i)
		}
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

// BenchmarkObjectInit benchmarks object initialization performance
func BenchmarkObjectInit(b *testing.B) {
	obj := objectivec.ObjectClass.Alloc()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		initialized := obj.Init()
		_ = initialized
	}
}

// BenchmarkObjectIsEqualTo benchmarks IsEqualTo performance
func BenchmarkObjectIsEqualTo(b *testing.B) {
	obj1 := objectivec.NewObject()
	obj2 := objectivec.NewObject()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equal := obj1.IsEqualTo(obj2)
		_ = equal
	}
}

// TestIObjectInterface verifies that Object implements IObject
func TestIObjectInterface(t *testing.T) {
	obj := objectivec.NewObject()

	// Verify Object can be assigned to IObject interface
	var iobj objectivec.IObject = obj

	// Verify the interface value is not nil
	if iobj == nil {
		t.Error("IObject interface should not be nil")
	}

	t.Logf("Object successfully implements IObject interface")
}
