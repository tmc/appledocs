// Package-level convenience helpers for Foundation.
//
// This file contains hand-written helpers that complement the generated bindings.
// It is not overwritten during code generation.
package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// String convenience methods

// StringWithGoString creates an NSString from a Go string.
// This is a convenience wrapper around stringWithUTF8String:.
//
// Example:
//
//	str := foundation.StringWithGoString("Hello, World!")
func StringWithGoString(s string) String {
	strClass := objc.GetClass("NSString")
	nsStr := objc.ID(strClass).Send(objc.RegisterName("stringWithUTF8String:"), s)
	return StringFrom(unsafe.Pointer(nsStr))
}

// GoString returns the Go string representation of an NSString.
// This is a convenience method that calls UTF8String and converts to Go string.
func (s String) GoString() string {
	// Get UTF8String (returns const char*)
	cStr := objc.Send[*byte](s.ID, objc.RegisterName("UTF8String"))
	if cStr == nil {
		return ""
	}
	// Convert C string to Go string using unsafe
	// Find string length
	length := 0
	for ptr := cStr; *ptr != 0; ptr = (*byte)(unsafe.Add(unsafe.Pointer(ptr), 1)) {
		length++
	}
	// Create Go string from C bytes
	return string(unsafe.Slice(cStr, length))
}

// Length returns the number of UTF-16 code units in the string.
func (s String) Length() uint {
	return objc.Send[uint](s.ID, objc.RegisterName("length"))
}
