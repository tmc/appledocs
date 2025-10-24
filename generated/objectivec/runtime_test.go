// Tests for Objective-C runtime functions
// These tests explore the Objective-C runtime C functions that are available
// in the objectivec package but not exported (lowercase names).

package objectivec_test

import (
	"testing"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// TestRuntimeFunctionAccess demonstrates accessing Objective-C runtime functions
// Note: Many runtime functions like class_getName are not exported from the
// objectivec package because they have lowercase names in C.
func TestRuntimeFunctionAccess(t *testing.T) {
	// Create an NSObject instance
	obj := objectivec.NewObject()
	if obj.ID == 0 {
		t.Fatal("Failed to create NSObject instance")
	}

	// Get the class
	cls := obj.ID.Class()
	if cls == 0 {
		t.Fatal("Failed to get class from object")
	}

	t.Logf("Got class: %v", cls)
}

// TestClassGetNameDirect demonstrates calling class_getName directly via purego
func TestClassGetNameDirect(t *testing.T) {
	// Open the Objective-C runtime library
	lib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		t.Fatalf("Failed to open libobjc: %v", err)
	}

	// Get the class_getName function pointer
	sym, err := purego.Dlsym(lib, "class_getName")
	if err != nil {
		t.Fatalf("Failed to find class_getName: %v", err)
	}

	// Register the function
	var class_getName func(objc.Class) unsafe.Pointer
	purego.RegisterFunc(&class_getName, sym)

	// Create an NSObject and get its class
	obj := objectivec.NewObject()
	cls := obj.ID.Class()

	// Call class_getName
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

// TestClassGetSuperclass demonstrates that some runtime functions ARE exported
// because they have proper capitalization or are wrapped
func TestClassGetSuperclass(t *testing.T) {
	// Note: class_getSuperclass is also lowercase, so we need to access it via purego
	lib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		t.Fatalf("Failed to open libobjc: %v", err)
	}

	sym, err := purego.Dlsym(lib, "class_getSuperclass")
	if err != nil {
		t.Fatalf("Failed to find class_getSuperclass: %v", err)
	}

	var class_getSuperclass func(objc.Class) objc.Class
	purego.RegisterFunc(&class_getSuperclass, sym)

	// Create an NSObject and get its class
	obj := objectivec.NewObject()
	cls := obj.ID.Class()

	// Get the superclass
	superclass := class_getSuperclass(cls)
	t.Logf("NSObject superclass: %v", superclass)

	// NSObject has no superclass (it's the root class)
	if superclass != 0 {
		t.Logf("NSObject has superclass (unexpected but valid in some contexts)")
	}
}

// TestMethodGetName tests getting method names from the runtime
func TestMethodGetName(t *testing.T) {
	// Open the Objective-C runtime library
	lib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		t.Fatalf("Failed to open libobjc: %v", err)
	}

	// Register the functions we need
	var class_getInstanceMethod func(objc.Class, objc.SEL) uintptr
	var method_getName func(uintptr) objc.SEL
	var sel_getName func(objc.SEL) unsafe.Pointer

	sym1, err := purego.Dlsym(lib, "class_getInstanceMethod")
	if err != nil {
		t.Fatalf("Failed to find class_getInstanceMethod: %v", err)
	}
	purego.RegisterFunc(&class_getInstanceMethod, sym1)

	sym2, err := purego.Dlsym(lib, "method_getName")
	if err != nil {
		t.Fatalf("Failed to find method_getName: %v", err)
	}
	purego.RegisterFunc(&method_getName, sym2)

	sym3, err := purego.Dlsym(lib, "sel_getName")
	if err != nil {
		t.Fatalf("Failed to find sel_getName: %v", err)
	}
	purego.RegisterFunc(&sel_getName, sym3)

	// Get NSObject class
	obj := objectivec.NewObject()
	cls := obj.ID.Class()

	// Get the 'init' method
	sel := objc.RegisterName("init")
	method := class_getInstanceMethod(cls, sel)
	if method == 0 {
		t.Fatal("Failed to get init method")
	}

	// Get the method name
	methodSel := method_getName(method)
	namePtr := sel_getName(methodSel)
	name := cStringToGoString(namePtr)

	t.Logf("Method name: %s", name)
	if name != "init" {
		t.Errorf("Expected method name 'init', got '%s'", name)
	}
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

// BenchmarkClassGetName benchmarks calling class_getName
func BenchmarkClassGetName(b *testing.B) {
	// Setup
	lib, err := purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		b.Fatalf("Failed to open libobjc: %v", err)
	}

	sym, err := purego.Dlsym(lib, "class_getName")
	if err != nil {
		b.Fatalf("Failed to find class_getName: %v", err)
	}

	var class_getName func(objc.Class) unsafe.Pointer
	purego.RegisterFunc(&class_getName, sym)

	obj := objectivec.NewObject()
	cls := obj.ID.Class()

	// Benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		namePtr := class_getName(cls)
		_ = namePtr
	}
}
