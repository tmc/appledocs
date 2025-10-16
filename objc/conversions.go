package objc

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
)

var (
	// Foundation framework handle
	foundationHandle uintptr

	// NSString selectors
	sel_stringWithUTF8String objc.SEL
	sel_UTF8String           objc.SEL
	sel_length               objc.SEL

	// NSString class
	nsStringClass objc.Class
)

func init() {
	// Load Foundation framework
	var err error
	foundationHandle, err = purego.Dlopen("/System/Library/Frameworks/Foundation.framework/Foundation", purego.RTLD_GLOBAL|purego.RTLD_NOW)
	if err != nil {
		panic(err)
	}

	// Cache common selectors
	sel_stringWithUTF8String = objc.RegisterName("stringWithUTF8String:")
	sel_UTF8String = objc.RegisterName("UTF8String")
	sel_length = objc.RegisterName("length")

	// Get NSString class
	nsStringClass = objc.GetClass("NSString")
	if nsStringClass == 0 {
		panic("NSString class not found")
	}
}

// ToNSString converts a Go string to an NSString (objc.ID).
//
// The returned NSString is autoreleased. The caller does not need to
// manually release it unless retained explicitly.
//
// Example:
//
//	nsStr := objc.ToNSString("Hello, World!")
//	defer nsStr.Send(objc.RegisterName("release"))
func ToNSString(s string) objc.ID {
	// Null-terminate the string
	cStr := append([]byte(s), 0)
	return objc.ID(nsStringClass).Send(sel_stringWithUTF8String, cStr)
}

// ToGoString converts an NSString (objc.ID) to a Go string.
//
// This function handles nil NSString objects by returning an empty string.
//
// Example:
//
//	goStr := objc.ToGoString(nsString)
func ToGoString(nsString objc.ID) string {
	if nsString == 0 {
		return ""
	}

	// Get UTF8String pointer
	cStrPtr := objc.Send[uintptr](nsString, sel_UTF8String)
	if cStrPtr == 0 {
		return ""
	}

	// Convert C string to Go string
	return cStringToGo(cStrPtr)
}

// cStringToGo converts a null-terminated C string to a Go string.
func cStringToGo(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}

	var bytes []byte
	for i := uintptr(0); ; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		if b == 0 {
			break
		}
		bytes = append(bytes, b)
	}
	return string(bytes)
}

// StringLength returns the length of an NSString.
//
// Example:
//
//	length := objc.StringLength(nsString)
func StringLength(nsString objc.ID) uint64 {
	if nsString == 0 {
		return 0
	}
	return objc.Send[uint64](nsString, sel_length)
}
