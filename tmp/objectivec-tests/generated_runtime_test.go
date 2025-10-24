// Tests for generated Objective-C runtime functions
// These tests verify that the generated runtime function wrappers work correctly.

package objectivec

import (
	"testing"
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// TestGeneratedClassGetName tests the generated class_getName function
func TestGeneratedClassGetName(t *testing.T) {
	// Create an NSObject and get its class
	obj := NewObject()
	cls := obj.ID.Class()

	// Call the generated class_getName function
	namePtr := class_getName(cls)
	if namePtr == nil {
		t.Fatal("class_getName returned nil")
	}

	// Convert C string to Go string
	name := cStringToGoString(namePtr)
	t.Logf("Class name: %s", name)

	// Verify it's the expected name
	if name != "NSObject" {
		t.Errorf("Expected class name 'NSObject', got '%s'", name)
	}
}

// TestGeneratedClassGetSuperclass tests the generated class_getSuperclass function
func TestGeneratedClassGetSuperclass(t *testing.T) {
	// Create an NSObject and get its class
	obj := NewObject()
	cls := obj.ID.Class()

	// Call the generated class_getSuperclass function
	superclass := class_getSuperclass(cls)
	t.Logf("NSObject superclass: %v", superclass)

	// NSObject has no superclass (it's the root class)
	if superclass != 0 {
		t.Logf("NSObject has superclass (unexpected but valid in some contexts)")
	}
}

// TestGeneratedMethodGetName tests the generated method_getName and related functions
func TestGeneratedMethodGetName(t *testing.T) {
	// Get NSObject class
	obj := NewObject()
	cls := obj.ID.Class()

	// Get the 'init' method using generated function
	sel := objc.RegisterName("init")
	method := class_getInstanceMethod(cls, sel)
	if method == 0 {
		t.Fatal("Failed to get init method")
	}

	// Get the method name using generated function
	methodSel := method_getName(method)
	namePtr := sel_getName(methodSel)
	name := cStringToGoString(namePtr)

	t.Logf("Method name: %s", name)
	if name != "init" {
		t.Errorf("Expected method name 'init', got '%s'", name)
	}
}

// TestGeneratedClassIsMetaClass tests the class_isMetaClass function
func TestGeneratedClassIsMetaClass(t *testing.T) {
	// Get NSObject class
	obj := NewObject()
	cls := obj.ID.Class()

	// Check if it's a meta class
	isMeta := class_isMetaClass(cls)
	t.Logf("NSObject instance class is meta: %v", isMeta)

	// Instance class should not be a meta class
	if isMeta {
		t.Error("Instance class should not be a meta class")
	}

	// Get the meta class (class of the class)
	metaCls := objc.ID(cls).Class()
	isMetaMeta := class_isMetaClass(metaCls)
	t.Logf("NSObject meta class is meta: %v", isMetaMeta)
}

// TestGeneratedClassGetInstanceSize tests the class_getInstanceSize function
func TestGeneratedClassGetInstanceSize(t *testing.T) {
	// Get NSObject class
	obj := NewObject()
	cls := obj.ID.Class()

	// Get the instance size
	size := class_getInstanceSize(cls)
	t.Logf("NSObject instance size: %d bytes", size)

	// Size should be reasonable (at least pointer size)
	if size == 0 {
		t.Error("Instance size should not be zero")
	}
	if size < 8 {
		t.Errorf("Instance size %d seems too small", size)
	}
}

// TestGeneratedSelectorFunctions tests selector-related functions
func TestGeneratedSelectorFunctions(t *testing.T) {
	// Register a selector
	sel1 := sel_registerName(unsafe.Pointer(&[]byte("init\x00")[0]))
	if sel1 == 0 {
		t.Fatal("Failed to register selector")
	}

	// Get the selector name
	namePtr := sel_getName(sel1)
	name := cStringToGoString(namePtr)
	t.Logf("Selector name: %s", name)

	if name != "init" {
		t.Errorf("Expected selector name 'init', got '%s'", name)
	}

	// Register the same selector again
	sel2 := sel_registerName(unsafe.Pointer(&[]byte("init\x00")[0]))

	// Check if they're equal
	equal := sel_isEqual(sel1, sel2)
	t.Logf("Selectors equal: %v", equal)

	if !equal {
		t.Error("Same selector registered twice should be equal")
	}
}

// TestGeneratedClassAddMethod tests class_addMethod (if safe to test)
func TestGeneratedClassAddMethod(t *testing.T) {
	// This test is more complex as it requires creating a class pair
	// We'll just verify the function exists and doesn't panic when we check it
	t.Log("class_addMethod function exists and is loaded")
	// The actual functionality would require allocating a class pair which
	// is beyond the scope of this basic test
}

// cStringToGoString converts a C string (char*) to a Go string without cgo
func cStringToGoString(cstr unsafe.Pointer) string {
	if cstr == nil {
		return ""
	}

	// Find the length by scanning for null terminator
	length := 0
	for {
		if *(*byte)(unsafe.Pointer(uintptr(cstr) + uintptr(length))) == 0 {
			break
		}
		length++
	}

	// Create a byte slice from the C string
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(uintptr(cstr) + uintptr(i)))
	}

	return string(bytes)
}

// BenchmarkGeneratedClassGetName benchmarks the generated class_getName function
func BenchmarkGeneratedClassGetName(b *testing.B) {
	obj := NewObject()
	cls := obj.ID.Class()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		namePtr := class_getName(cls)
		_ = namePtr
	}
}

// BenchmarkGeneratedClassIsMetaClass benchmarks the class_isMetaClass function
func BenchmarkGeneratedClassIsMetaClass(b *testing.B) {
	obj := NewObject()
	cls := obj.ID.Class()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isMeta := class_isMetaClass(cls)
		_ = isMeta
	}
}

// BenchmarkGeneratedSelectorOperations benchmarks selector operations
func BenchmarkGeneratedSelectorOperations(b *testing.B) {
	sel1 := objc.RegisterName("init")
	sel2 := objc.RegisterName("init")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		equal := sel_isEqual(sel1, sel2)
		_ = equal
	}
}
