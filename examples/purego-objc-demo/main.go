// Package main demonstrates using github.com/ebitengine/purego/objc directly
// to interact with Objective-C frameworks without a custom runtime layer.
package main

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

func main() {
	// Load Foundation framework
	_, err := purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_GLOBAL|purego.RTLD_NOW)
	if err != nil {
		panic(err)
	}

	// Get NSString class
	nsStringClass := objc.GetClass("NSString")
	if nsStringClass == 0 {
		panic("NSString class not found")
	}

	// Cache selectors (RegisterName grabs global lock)
	sel_stringWithUTF8String := objc.RegisterName("stringWithUTF8String:")
	sel_UTF8String := objc.RegisterName("UTF8String")
	sel_length := objc.RegisterName("length")
	sel_uppercaseString := objc.RegisterName("uppercaseString")

	// Create NSString using class method
	goString := "Hello from purego/objc!\x00" // null-terminated
	nsString := objc.ID(nsStringClass).Send(sel_stringWithUTF8String, goString)

	// Get length using generic Send[T]
	length := objc.Send[uint64](nsString, sel_length)
	fmt.Printf("String length: %d\n", length)

	// Convert to uppercase
	upperString := nsString.Send(sel_uppercaseString)

	// Convert back to Go string (simplified - would need proper conversion)
	upperCString := objc.Send[uintptr](upperString, sel_UTF8String)
	fmt.Printf("Uppercase: %s\n", cStringToGo(upperCString))
	fmt.Printf("Original:  %s\n", goString)
}

// cStringToGo converts a C string pointer to a Go string
func cStringToGo(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var s []byte
	for i := uintptr(0); ; i++ {
		c := *(*byte)(unsafe.Pointer(ptr + i))
		if c == 0 {
			break
		}
		s = append(s, c)
	}
	return string(s)
}
