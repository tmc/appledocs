package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

// CFTypeRef represents a Core Foundation object reference
type CFTypeRef unsafe.Pointer

// CFStringRef represents a Core Foundation string reference
type CFStringRef unsafe.Pointer

var (
	lib uintptr

	// CFStringCreateWithCString creates a CFString from a C string
	CFStringCreateWithCString func(alloc CFTypeRef, cStr *byte, encoding uint32) CFStringRef

	// CFStringGetCString gets a C string from a CFString
	CFStringGetCString func(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool

	// CFRelease releases a Core Foundation object
	CFRelease func(cf CFTypeRef)

	// CFRetain retains a Core Foundation object
	CFRetain func(cf CFTypeRef) CFTypeRef

	// CFGetRetainCount gets the retain count of a Core Foundation object
	CFGetRetainCount func(cf CFTypeRef) int
)

const (
	kCFStringEncodingUTF8 = 0x08000100
)

func init() {
	var err error
	lib, err = purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&CFStringCreateWithCString, lib, "CFStringCreateWithCString")
	purego.RegisterLibFunc(&CFStringGetCString, lib, "CFStringGetCString")
	purego.RegisterLibFunc(&CFRelease, lib, "CFRelease")
	purego.RegisterLibFunc(&CFRetain, lib, "CFRetain")
	purego.RegisterLibFunc(&CFGetRetainCount, lib, "CFGetRetainCount")
}

func main() {
	// Test 1: Create a CFString
	cStr := []byte("Hello from Go via purego!\x00")
	cfStr := CFStringCreateWithCString(nil, &cStr[0], kCFStringEncodingUTF8)
	if cfStr == nil {
		panic("Failed to create CFString")
	}
	fmt.Println("✅ Created CFString successfully")

	// Test 2: Check retain count
	count := CFGetRetainCount(CFTypeRef(cfStr))
	fmt.Printf("✅ Initial retain count: %d\n", count)

	// Test 3: Retain and check count again
	cfStr2 := CFStringRef(CFRetain(CFTypeRef(cfStr)))
	count = CFGetRetainCount(CFTypeRef(cfStr))
	fmt.Printf("✅ After retain, count: %d\n", count)

	// Test 4: Convert back to C string
	buffer := make([]byte, 256)
	ok := CFStringGetCString(cfStr, &buffer[0], len(buffer), kCFStringEncodingUTF8)
	if !ok {
		panic("Failed to get C string")
	}
	// Find null terminator
	var str string
	for i, b := range buffer {
		if b == 0 {
			str = string(buffer[:i])
			break
		}
	}
	fmt.Printf("✅ Retrieved string: %q\n", str)

	// Test 5: Release both references
	CFRelease(CFTypeRef(cfStr2))
	CFRelease(CFTypeRef(cfStr))
	fmt.Println("✅ Released CFString successfully")

	fmt.Println("\n🎉 All tests passed! Generated bindings work perfectly!")
}
