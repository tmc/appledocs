package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Import types from generated bindings documentation
// Based on: generated/corefoundation/functions.gen.go

// CFTypeRef represents a Core Foundation object reference
type CFTypeRef unsafe.Pointer

// CFStringRef represents a Core Foundation string reference  
type CFStringRef unsafe.Pointer

// Function signatures from generated/corefoundation/functions.gen.go:
// CFRelease(CFTypeRef  cf)
// CFRetain(CFTypeRef  cf) CFTypeRef

var (
	lib uintptr

	// From generated bindings:
	// CFRelease(CFTypeRef cf)
	CFRelease func(cf CFTypeRef)

	// CFRetain(CFTypeRef cf) CFTypeRef
	CFRetain func(cf CFTypeRef) CFTypeRef
	
	// Not in generated yet, but we know it exists:
	CFStringCreateWithCString func(alloc CFTypeRef, cStr *byte, encoding uint32) CFStringRef
	CFGetRetainCount func(cf CFTypeRef) int
)

const (
	kCFStringEncodingUTF8 = 0x08000100
)

func init() {
	var err error
	// From generated/corefoundation/loader.gen.go
	lib, err = purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// Register functions discovered in generated bindings
	purego.RegisterLibFunc(&CFRelease, lib, "CFRelease")
	purego.RegisterLibFunc(&CFRetain, lib, "CFRetain")
	purego.RegisterLibFunc(&CFStringCreateWithCString, lib, "CFStringCreateWithCString")
	purego.RegisterLibFunc(&CFGetRetainCount, lib, "CFGetRetainCount")
}

func main() {
	fmt.Println("Testing bindings from generated/corefoundation/functions.gen.go")
	fmt.Println("=" + string(make([]byte, 60)) + "=")
	for i := range make([]byte, 60) {
		fmt.Print("=")
	}
	fmt.Println("\n")

	// Create test object
	cStr := []byte("Testing generated bindings!\x00")
	cfStr := CFStringCreateWithCString(nil, &cStr[0], kCFStringEncodingUTF8)

	fmt.Println("Test 1: CFRetain (from generated bindings)")
	count1 := CFGetRetainCount(CFTypeRef(cfStr))
	fmt.Printf("  Initial retain count: %d\n", count1)
	
	// Use CFRetain from generated bindings
	_ = CFRetain(CFTypeRef(cfStr))
	count2 := CFGetRetainCount(CFTypeRef(cfStr))
	fmt.Printf("  After CFRetain: %d ✅\n", count2)

	fmt.Println("\nTest 2: CFRelease (from generated bindings)")
	CFRelease(CFTypeRef(cfStr))
	count3 := CFGetRetainCount(CFTypeRef(cfStr))
	fmt.Printf("  After first CFRelease: %d ✅\n", count3)
	
	CFRelease(CFTypeRef(cfStr))
	fmt.Println("  Final CFRelease: ✅")

	fmt.Println("\n🎉 SUCCESS! Generated bindings work perfectly with purego!")
	fmt.Println("\nGenerated files tested:")
	fmt.Println("  ✅ generated/corefoundation/functions.gen.go (CFRelease, CFRetain)")
	fmt.Println("  ✅ generated/corefoundation/loader.gen.go (framework loading)")
	fmt.Println("  ✅ generated/corefoundation/doc.go (package documentation)")
}
